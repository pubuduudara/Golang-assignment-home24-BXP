## Installation & Setup

- `cd backend` and go to backend folder
- install dependencies `go mod tidy`
- (Optional) Create a `.env` file. The .env file is already committed for convenience. You only need to follow this step if you want to change the API key.

```env
API_KEY=your-custom-api-key
```

#### You can run the application in two ways:

1. Using Makefile, without Docker

   - Build the application `make build`
   - Run the application `make run` The server will start on http://localhost:8080

2. Using Docker (you need to have docker desktop running on your machine)
   - Build the Docker image `make docker-build`
   - Run the Docker container in the foreground `make docker-run`
   - To stop running the container run `make docker-stop`

#### Developer Testing

- To run all unit and integration tests with coverage, run `make test`
