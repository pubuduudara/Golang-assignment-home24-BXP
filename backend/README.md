## Installation & Setup

- `cd backend` and go to backend folder
- install dependencies `go mod tidy`
- (Optional) Create a `.env` file. The .env file is already committed for convenience. You only need to follow this step if you want to change the API key.

```env
API_KEY=your-custom-api-key
```

#### You can run the application in two ways. Once below is completed, backend app is running sucessfully:

1. Using Makefile, without Docker (Recommended)

   - Build the application `make build`
   - Run the application `make run`. The server will start on http://localhost:8080

2. Using Docker (you need to have docker desktop running on your machine)
   - Build the Docker image `make docker-build`
   - Run the Docker container in the foreground `make docker-run`
   - To stop running the container run `make docker-stop`

#### Developer Testing

- To run all unit and integration tests with coverage, run `make test`
- You can see the test coverage in the terminal

## Tools and Technologies

1. Backend
   - Golang – Core application logic
   - Chi – HTTP router (Lightweight router)
   - Slog – Structured logging
   - Net/HTML – HTML parsing
   - Godotenv – Loads environment variables from `.env` file into the app
   - Docker – Containerization
   - Makefile – Build and task automation
2. Frontend
   - React with Vite
   - Axios – For making HTTP requests to the backend API

## Design Considerations

- Project is organized with a clear separation of concerns,
  - `internal/` is used for private application logic to prevent unintended imports from outside the module, following Go best practices.
  - `internal/api` holds HTTP routing and handlers.
  - `internal/services` includes core business logic for analyzing web pages.
  - `internal/middleware` contains reusable middleware (e.g., API key auth).
  - `internal/utils` provides helper functions and centralized logging setup.
- Go concurrency is used to check the accessibility of links in parallel using goroutines, wait groups, and a semaphore to limit concurrency and avoid overloading external websites.
- Unit and integration tests are implemented to cover business logic, API routes, and middleware behavior, without relying on real external HTTP calls.

## Features

- API security is enforced using an API key, which must be included in the `X-API-Key` request header. This is validated by middleware to ensure secure access. The Frontend includes this key in the header when making requests.
- The API maintains a consistent JSON response structure, making it easier for frontend developers to handle and parse responses. Also it includes a `status`, which will be a boolean which represents whether the API call was a sucess or a failure:
  ```json
  {
    "status": true | false,
    "data": <result or error message>
  }
  ```
- Environment variables are used for config like the API key, with .env file support for local development.
- The app is containerized with Docker and all build/run/test tasks are automated using a Makefile
- Uses Go's log/slog package with a custom wrapper for consistent, leveled logging (Info, Warn, Error).
- Error responses are automatically logged and returned in a standardized JSON format, including the appropriate HTTP status code and message, helping both frontend and backend developers debug efficiently

## Possible Improvements

- Add more unit test cases to increase the overall coverage (current coverage is ~75%)
- Add support for detecting more HTML versions beyond HTML5 by parsing complete DOCTYPE metadata.
- Introduce Swagger (OpenAPI) documentation for better API usability, testing and documentation.
- Add rate limiting middleware to prevent abuse of the API.
- Introduce structured error types with custom codes for better error classification
- Implement request/response logging with trace IDs for better debugging
- Improve the frontend UI
