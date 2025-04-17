# Frontend – Web Page Analyzer

## Installation & Setup

### 1. Navigate to the frontend folder

```bash
cd frontend
```

### 2. Install dependencies

```bash
npm install
```

### 3. (Optional) Create a `.env` file

> The `.env` file is already committed to the repo for convenience.  
> You only need to follow this step **if you want to override the default API key**.

```env
VITE_API_KEY=your-custom-api-key
VITE_API_URL=http://localhost:8080
```

---

### 4. Start the development server

```bash
npm run dev
```

The app will be available at:

```
http://localhost:5173
```

You can load the UI from above URL.

---

### Usage Instructions

Before using the app, make sure to **run the backend Go server** as instructed in the `backend/README.md`.

Once the frontend is running at `http://localhost:5173`:

1. Enter the URL you want to analyze into the input field.
2. Click the **Analyze** button.
3. The app will send a request to the backend and display the results.

---
