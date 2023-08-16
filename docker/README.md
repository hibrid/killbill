:toc: macro
:toc-title:
:toclevels: 9

[[docker-compose-recipes]]
# docker-compose recipes

toc::[]

## Quick start

https://docs.docker.com/compose/[Docker Compose] is the easiest way to set up the Kill Bill stack:

```
docker-compose -f docker-compose.kb.yml -p kb up
```

Note: if you are using Docker Machine, use `docker-machine env <name>` or the environment variable `$DOCKER_HOST` to get the IP address of the container.

Then go to `http://<IP_ADDRESS>:9090/` and log-in via `admin/password`.

In the rest of this documentation, we will assumed `IP_ADDRESS` is 127.0.0.1.

## Logging and Metrics infrastructure

[[logging]]
### Logging

Refer to the https://www.elastic.co/guide/index.html[Elastic Stack] documentation for how to set up the ELK stack.

Here's a quick tutorial to get you started:

* Clone https://github.com/deviantony/docker-elk
* Update `logstash/pipeline/logstash.conf` and enable the Logstash JSON codec:
```
diff --git a/logstash/pipeline/logstash.conf b/logstash/pipeline/logstash.conf
index 40ca757..6ccbfb2 100644
--- a/logstash/pipeline/logstash.conf
+++ b/logstash/pipeline/logstash.conf
@@ -5,6 +5,7 @@ input {

        tcp {
                port => 5000
+               codec => json
        }
 }
```
* Run `docker-compose up` to start the stack. This takes a while. You can verify the stack is up by going to http://127.0.0.1:9200/ (user: elastic, password: changeme)
* Once the stack is up, redirect the Kill Bill container logs to Logstash using Logspout:
```
# Note: host.docker.internal is when using Docker Desktop for Mac. Replace with the internal IP address used by the host on other platforms
docker run --name="logspout" \
           --volume=/var/run/docker.sock:/var/run/docker.sock \
           gliderlabs/logspout \
           'tcp://host.docker.internal:5000?filter.name=kb_killbill_1'
```
* Go to Kibana at http://127.0.0.1:5601/ (user: elastic, password: changeme)
* Create an index pattern by going to http://127.0.0.1:5601/app/management/kibana/indexPatterns/create: enter `logstash-*` in the first screen and select `@timestamp` as the time filter field
* Go to `http://127.0.0.1:5601/app/discover` to visualize the Kill Bill logs

Notes:

* You need Docker Engine version 17.05 or newer and Docker Compose version 1.20.0 or newer
* You need at least 4GB of RAM allocated to Docker to run the entire stack

[[monitoring]]
### Metrics

Refer to the InfluxDB and Grafana documentation for how to set up the TIG stack.

Here's a quick tutorial to get you started:

* Run `docker-compose -f docker-compose.gi.yml -p gi up`
* Go to http://127.0.0.1:3000/ (user: admin, password: admin)
* Go to http://127.0.0.1:3000/datasources, click Add data source, select InfluxDB, enter:
  ** http://influxdb:8086 as the URL
  ** Enable Basic Auth (User: killbill, Password: killbill)
  ** Database name: killbill
* Click Save & Test
* Update `docker-compose.kb.yml`, set `KILLBILL_METRICS_INFLUXDB=true`, and recreate the Kill Bill docker stack

Note: you might also have to update `KILLBILL_METRICS_INFLUXDB_HOST` in `docker-compose.kb.yml`. The default value of `host.docker.internal` works when using Docker Desktop for Mac. Replace with the internal IP address used by the host on other platforms.