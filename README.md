# redis-probe
A simple Redis client application that tests the connectivity with a Redis server through a few REST API endpoints.

This Golang-based application runs as an http server that listens to requests from a configurable port and initiates a Redis client to connect to a configurable Redis server.

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

---

## Run redis-probe as a container

A sample container image containing the latest version of the `redis-probe` application is available from project `silval/redis-probe` in [Docker Hub](https://hub.docker.com/repository/docker/silval/redis-probe). The image is created with the [Dockerfile](./docker/Dockerfile) available in this repository.

A [sample deployment file](./deploy/redis-probe.yml) is provided to deploy `redis-probe` to a Kubernetes cluster. 

The sample deployment file deploys the container to a `dev` namespace, so make sure to create it or update that entry accordingly for your environment.

To create the dev namespace:
```
   kubectl create namespace dev 
```

Create a secret for the redis password in the same namespace:
```
   kubectl create secret generic redis-server --from-literal=password=REDIS_PASSWORD_GOES_HERE -n dev
```

Update the value of the environment variables in [`redis-probe.yml`](./deploy/redis-probe.yml) to configure the `redis-probe` to connect to yor Redis server. e.g. `REDIS_SERVER_ADDRESS`

Then deploy the application to your kubernetes cluster:
```
   kubectl apply -f deploy/redis-probe.yml
```

Get the IP address allocated to the `redis-probe` service:
```
   kubectl get service redis-probe -n dev
```

Access the `redis-probe` endpoints to test the connectivity to the Redis server:
```
   curl http://IP-ADDRESS-FROM-PREVIOUS-STEP

   curl http://IP-ADDRESS-FROM-PREVIOUS-STEP/ping

   curl http://IP-ADDRESS-FROM-PREVIOUS-STEP/set

   curl http://IP-ADDRESS-FROM-PREVIOUS-STEP/get

   curl http://IP-ADDRESS-FROM-PREVIOUS-STEP/delete
```

---