### ENV Variables - Client
```ini
CLIENTS=100
TIMEOUT=10m
WAIT=0s
RIG_HOST=localhost
GOMAXPROCS=4
```

## Run 1

```bash
cd src/run1
docker-compose up
```

## Run 2


```bash
cd src/run2
docker-compose up
```

## Run 6

### Start RIG

```bash
# Start Kafka first
export LOG_LEVEL=error
export KAFKA_BROKERS=localhost:9092
mix phx.server
```

### Start clients

```bash
CLIENTS=1000 TIMEOUT=1h go run src/run6/client.go
```

### Start loader

```bash
cd src/run6
./start
```
