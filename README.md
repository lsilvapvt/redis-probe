# redis-probe
A simple Golang application that exposes a REST API to test the connectivity with a Redis server.

The application runs as an http server that listens to requrests from a configurable port and initiates a Redis client to connect to a configurable Redis server.

Four endpoints are provided:
- /ping - sends a PING to Redis server
- /set  - saves the current time as the value for a MYTIME on the server
- /get  - retrieves the value of the MYTIME variable stored on the server
- /delete  - deletes the MYTIME variable stored on the server

## Usage

The following environment variables need to be configured in the environment or container in which the redis-probe will be executed on:
```
REDIS_SERVER_ADDRESS      # IP address or FQDN of the targeted Redis server. (e.g. 10.10.120.12)
REDIS_SERVER_PORT         # Port number for the Redis server (default is 6379)
REDIS_PASSWORD            # Password for the connection with the Redis server
REDIS_DATABASE            # Redis database instance to use (default is 0)
REDIS_PROBE_SERVICE_PORT  # Port number for the redis-probe to listen to for http requests (default is 8888)
```

Once the environment variables are configured, run the binary for the redis-probe and then you can access one of the endpoints listed above for the application.

```

$> curl http://10.10.120.12:8888/ping
PONG

```
## Run as a container in Kubernetes


