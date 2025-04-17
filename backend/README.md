# Backend – Web Page Analyzer (Golang)

## Installation & Setup

### 1. Navigate to the backend folder

```bash
cd backend
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. (Optional) Create a `.env` file

> The `.env` file is already committed for convenience.  
> You only need to follow this step **if you want to change the API key**.

```env
API_KEY=your-custom-api-key
```

---

### 4. Run the server

```bash
go run cmd/server/main.go
```

The backend will start at:

```
http://localhost:8080
```
