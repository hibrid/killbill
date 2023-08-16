
# Kill Bill with ELK Stack and Metrics Setup

This repository provides a Makefile to automate the setup of Kill Bill with the ELK stack for logging and TIG stack for metrics.

## Prerequisites

1. Ensure you have Docker and Docker Compose installed on your machine.
2. Ensure you have `git` installed for cloning repositories.

## How to Use

### 1. Quick Start

To quickly start the Kill Bill stack:

```
make quick-start
```

### 2. Logging with ELK

To set up the ELK stack for logging:

```
make setup-elk
```

### 3. Redirect Logs to Logstash

To redirect Kill Bill logs to Logstash using Logspout:

```
make log-redirect
```

Note: This command is platform-aware and will adjust based on whether you are running on macOS, WSL, or another UNIX-based system.

### 4. Metrics with TIG

To configure the Kill Bill stack for metrics with InfluxDB:

```
make configure-metrics
```

### 5. Full Setup

To execute all the above steps in sequence:

```
make all
```

This will provide you with access information and credentials for each system at the end.

### 6. Cleanup

To clean up any conflicting containers:

```
make clean
```

## Access Information and Credentials

After running `make all`, the following systems will be available:

- **Kill Bill UI**: http://127.0.0.1:9090/ - login: admin/password
- **Elasticsearch**: http://127.0.0.1:9200/ - login: elastic/changeme
- **Kibana**: http://127.0.0.1:5601/ - login: elastic/changeme
- **Grafana**: http://127.0.0.1:3000/ - login: admin/admin
- **InfluxDB**: http://influxdb:8086/ - login: killbill/killbill

Enjoy using Kill Bill with enhanced logging and metrics!



https://github.com/killbill/killbill-docs/tree/v3/catalogs

https://github.com/killbill/killbill-cloud/tree/master