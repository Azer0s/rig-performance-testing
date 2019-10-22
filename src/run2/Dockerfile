FROM golang:latest
COPY .. .
RUN go get github.com/r3labs/sse 
CMD ["go", "run", "run2/client.go"]
