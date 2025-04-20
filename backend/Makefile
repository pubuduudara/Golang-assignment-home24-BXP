.PHONY: build run test docker-build docker-run clean 

APP_NAME=web-analyzer
BUILD_DIR=build
DOCKER_IMAGE=web-analyzer
PORT=8080

# Build Go binary into build/
build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) ./cmd/server

run:
	go run ./cmd/server

test:
	go test ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out

docker-build:
	docker build -t $(DOCKER_IMAGE) .

docker-run:
	docker run --rm -p $(PORT):8080 $(DOCKER_IMAGE)

clean:
	rm -rf $(BUILD_DIR)
