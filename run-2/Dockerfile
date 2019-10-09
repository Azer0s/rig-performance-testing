FROM golang:latest
COPY . .
RUN go get github.com/r3labs/sse 
CMD ["go", "run", "src/client.go"]
