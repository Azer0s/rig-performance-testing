## Run 1

Send one message, send 1M ignored messages, send one message again.

### ENV Variables - Client
```ini
CLIENTS=1
TIMEOUT=10m
WAIT=30s
RIG_HOST=rig
```

### Start
```bash
cd src/run1
docker-compose up
```

## Run 2

Send 100k messages to 100 clients.

### ENV Variables - Client
```ini
CLIENTS=100
TIMEOUT=30m
WAIT=30s
RIG_HOST=rig
```

### Start
```bash
cd src/run2
docker-compose up
```

## Run 6

Send 1000 messages in 100 different event types to 1000 clients.

### ENV Variables - Client
```ini
CLIENTS=1000
TIMEOUT=1h
WAIT=30s
RIG_HOST=rig
```

### Start
```bash
cd src/run6
docker-compose up
```

