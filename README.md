# Gymshark Pack Calculation Service

This project is a sample Go application demonstrating how to fulfill customer orders using predefined pack sizes while minimizing overhead and the number of packs.

## Overview

Customers can order any number of items. The service determines the optimal combination of packs to:
1. Use only whole packs (no partial packs).
2. Ensure no more items than necessary are sent (minimal overhead).
3. Use the fewest packs possible given the above constraints.

By default, it uses pack sizes: `250, 500, 1000, 2000, 5000`.

**Features**:
- A `/calculate?items=<N>` endpoint that returns the optimal pack combination in JSON.
- A `/pack-sizes` endpoint to view, add, and delete pack sizes at runtime without code changes.
- A simple UI served at `/` to interact with the API from a browser.
- Easily configurable and ready for deployment in different environments (local, EC2, Lambda).

## Project Structure

```
.  
├── cmd  
│   └── server  
│       └── main.go // Entry point for running the server  
├── internal  
│   ├── config // Configuration loading  
│   ├── env // Environment loading (.env)  
│   ├── handlers // HTTP handlers for /calculate, /pack-sizes, and UI  
│   ├── storage // In-memory storage for pack sizes  
│   └── usecase // Business logic (DP calculation)  
├── static  
│   └── index.html // The UI file served at "/"  
└── test  
    └── calculation_test.go // Basic tests for calculation logic
```

- **`cmd/server`**: Runs the HTTP server.
- **`internal/usecase`**: Core calculation logic (DP algorithm).
- **`internal/handlers`**: Exposes HTTP endpoints for calculating packs, managing pack sizes, and serving UI.
- **`internal/storage`**: In-memory store for pack sizes, thread-safe.
- **`static/index.html`**: A simple HTML/JS UI to interact with the API.

## Requirements

- Go 1.20+ (or a recent version)
- Optional: `git` for cloning, `curl` for testing endpoints, `AWS CLI` for deployment tasks.
- Docker for containerization (optional).

## Getting Started Locally

1. **Clone the repository**:
   ```bash
   git clone https://github.com/chiranshu11/gymshark.git
   cd gymshark
   ```

2. **Initialize and tidy up modules:**
   ```bash
   go mod tidy
   ```

3. **Build and Run**:
   ```bash
   go build -o server ./cmd/server
   ./server
   ```

## Open in Browser

**Navigate to `http://localhost:8080/`:**

- The UI page allows you to:
  - View, add, and delete pack sizes.
  - Perform a calculation for a given number of items.

## API Endpoints

- **GET /pack-sizes**: Returns a JSON array of current pack sizes.
- **POST /pack-sizes** with `{ "size": <int> }`: Adds a new pack size.
- **DELETE /pack-sizes?size=<int>**: Removes a pack size.
- **GET /calculate?items=<N>**: Returns the calculation result as JSON.

## Example Usage

**Calculate packs for an order of 501 items:**

```bash
curl "http://localhost:8080/calculate?items=501"
```

**Response (example):**

```json
{
  "requested": 501,
  "used_sum": 750,
  "overhead": 249,
  "packs": {
    "250": 1,
    "500": 1
  }
}
```

---

### Updating Pack Sizes

**Add a new pack size (e.g., 750):**

```bash
curl -X POST -H "Content-Type: application/json" -d '{"size":750}' http://localhost:8080/pack-sizes
```

**Delete a pack size (e.g., 1000):**

```bash
curl -X DELETE "http://localhost:8080/pack-sizes?size=1000"
```

These changes take effect immediately and will be used in subsequent calculations.

## Docker Support

To run the service using Docker:

1. **Build the Docker image:**
   ```bash
   docker build -t gymshark-pack-service .
   ```

2. **Run the Docker container:**
   ```bash
   docker run -p 8080:8080 gymshark-pack-service
   ```

3. **Access the service:**
   - API: Navigate to `http://localhost:8080`
   - UI: Navigate to `http://localhost:8080/`