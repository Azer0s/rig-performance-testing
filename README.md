## Run 1

### Start RIG

```bash
# Start Kafka first
export LOG_LEVEL=error
export KAFKA_BROKERS=localhost:9092
mix phx.server
```

### Start clients

```bash
CLIENTS=1 GOMAXPROCS=2 go run run-1/src/client.go
```

### Start loader

```bash
python3 run-1/src/loader.py
```

## Run 2

### Start RIG

```bash
# Start Kafka first
export LOG_LEVEL=error
export KAFKA_BROKERS=localhost:9092
mix phx.server
```

### Start clients

```bash
CLIENTS=3500 GOMAXPROCS=4 go run run-1/src/client.go
```

### Start loader

```bash
./run-2/start
```

