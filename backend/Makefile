APP_NAME=web-analyzer
BUILD_DIR=build
DOCKER_IMAGE=web-analyzer
PORT=8080

# Build Go binary into build/
build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) ./cmd/server

run:
	source .env && go run ./cmd/server

test:
	go test ./... -cover

docker-build:
	docker build -t $(DOCKER_IMAGE) .

docker-run:
	docker run --rm -p $(PORT):8080 --env-file .env $(DOCKER_IMAGE)

docker-up:
	docker run -d -p $(PORT):8080 --env-file .env --name $(APP_NAME) $(DOCKER_IMAGE)

docker-stop:
	docker stop $(APP_NAME)

clean:
	rm -rf $(BUILD_DIR)
