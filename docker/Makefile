
# Makefile to set up Kill Bill with ELK stack and Metrics

# Quick start Kill Bill
quick-start:
	docker-compose -f docker-compose.kb.yml -p kb up -d

# Setup ELK for logging
setup-elk:
	@if [ -d "docker-elk" ]; then \
		echo "docker-elk directory exists. Updating repository..."; \
		cd docker-elk && git pull; \
	else \
		git clone https://github.com/deviantony/docker-elk; \
		sed -i 's/tcp {/tcp {\n\t\tcodec => json/g' docker-elk/logstash/pipeline/logstash.conf; \
	fi
	docker-compose -f docker-elk/docker-compose.yml up setup
	docker-compose -f docker-elk/docker-compose.yml up --build -d

# Redirect Kill Bill logs to Logstash using Logspout
log-redirect:
	@docker rm -f logspout || true
	@if cat /proc/version | grep microsoft; then \
		docker run  -itd --network=docker-elk_elk --name="logspout"  \
			--volume=/var/run/docker.sock:/var/run/docker.sock \
			gliderlabs/logspout \
			'tcp://docker-elk_logstash_1:50000?filter.name=kb_killbill_1'; \
	elif [ "$$(uname)" = "Darwin" ]; then \
		docker run -itd --network=docker-elk_elk --name="logspout" -d \
			--volume=/var/run/docker.sock:/var/run/docker.sock \
			gliderlabs/logspout \
			'tcp://host.docker.internal:50000?filter.name=kb_killbill_1'; \
	else \
		DOCKER_HOST_IP=$$(ip route | grep default | awk '{print $$3}'); \
		docker run -itd --network=docker-elk_elk --name="logspout" -d \
			--volume=/var/run/docker.sock:/var/run/docker.sock \
			gliderlabs/logspout \
			"tcp://$$DOCKER_HOST_IP:50000?filter.name=kb_killbill_1"; \
	fi

clean:
	docker stop $$(docker ps -a -q --filter name=killbill) || true
	docker rm $$(docker ps -a -q --filter name=killbill) || true
	docker stop $$(docker ps -a -q --filter name=kaui) || true
	docker rm $$(docker ps -a -q --filter name=kaui) || true
	docker stop $$(docker ps -a -q --filter name=db) || true
	docker rm $$(docker ps -a -q --filter name=db) || true
	docker stop $$(docker ps -a -q --filter name=influxdb) || true
	docker rm $$(docker ps -a -q --filter name=influxdb) || true
	docker-compose -f docker-elk/docker-compose.yml down || true
	docker-compose -f docker-compose.kb.yml -p kb down || true
	docker stop $$(docker ps -a -q --filter name=logspout) || true
	docker rm $$(docker ps -a -q --filter name=logspout) || true



configure-metrics:
	sed -i 's/KILLBILL_METRICS_INFLUXDB=false/KILLBILL_METRICS_INFLUXDB=true/g' docker-compose.kb.yml
	docker-compose -f docker-compose.kb.yml -p kb up -d

all: quick-start setup-elk log-redirect configure-metrics


# Print access information and credentials
info:
	@echo "Kill Bill UI: http://127.0.0.1:9090/ - login: admin/password"
	@echo "Elasticsearch: http://127.0.0.1:9200/ - login: elastic/changeme"
	@echo "Kibana: http://127.0.0.1:5601/ - login: elastic/changeme"
	@echo "Grafana: http://127.0.0.1:3000/ - login: admin/admin"
	@echo "InfluxDB: http://influxdb:8086/ - login: killbill/killbill"

all: quick-start setup-elk log-redirect configure-metrics info
