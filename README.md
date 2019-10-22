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
CLIENTS=1 GOMAXPROCS=2 go run src/run1/client.go
```

### Start loader

```bash
cd src/run1
python3 loader.py
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
CLIENTS=3500 GOMAXPROCS=4 go run run2/src/client.go
```

### Start loader

```bash
cd src/run2
./start
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
CLIENTS=1000 GOMAXPROCS=4 go run src/run6/client.go
```

### Start loader

```bash
cd src/run6
./start
```
