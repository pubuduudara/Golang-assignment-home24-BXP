FROM golang:1.24 as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

# Build to /app/build
RUN mkdir -p build && go build -o build/web-analyzer ./cmd/server

FROM debian:bookworm-slim

WORKDIR /app

# install CA certificates (fixes TLS trust issues)
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*


# Copy built binary and .env
COPY --from=builder /app/build/web-analyzer .
COPY .env .

EXPOSE 8080

CMD ["./web-analyzer"]

