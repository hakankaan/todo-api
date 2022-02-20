##
## Build
##
FROM golang:1.17-buster AS builder


WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

WORKDIR /app/cmd/server

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o todo-api
RUN chmod +x todo-api

##
## Deploy
##
FROM debian:buster-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/cmd/server/todo-api /app/cmd/server/todo-api

# Run the web service on container startup.
CMD ["/app/cmd/server/todo-api"]