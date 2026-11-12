# OLX API

A lightweight, high-performance REST API service built with Go (Golang) for the OLX platform. This project provides the foundational infrastructure for building scalable classified advertisement services.

**Live Demo**: [https://olx-api-wzl0.onrender.com/healthz](https://olx-api-wzl0.onrender.com/healthz)

**Developer**: Dinesh Prajapati

---

## Table of Contents

- [Features](#features)
- [Tech Stack](#tech-stack)
- [Folder Structure](#folder-structure)
- [API Endpoints](#api-endpoints)
- [Setup Guide](#setup-guide)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Environment Configuration](#environment-configuration)
  - [Running the Application](#running-the-application)
- [Usage](#usage)
- [Project Structure](#project-structure)

---

## Features

- **Health Check Endpoint**: Monitor application status and availability
- **Home Endpoint**: Welcome message and API information
- **Environment Configuration**: Secure environment variable management
- **Database Connectivity**: PostgreSQL database integration with connection pooling
- **Production-Ready**: Configured with proper timeouts for production deployment
- **Lightweight**: Minimal dependencies for fast startup and low memory footprint

---

## Tech Stack

| Category | Technology | Version |
|----------|------------|---------|
| **Language** | Go (Golang) | 1.26.7 |
| **Framework** | Standard Library (net/http) | - |
| **Database** | PostgreSQL | - |
| **Dependency Management** | Go Modules | - |
| **Environment Management** | godotenv | v1.5.1 |
| **Hosting** | Render | - |

---

## Folder Structure

```
olx-api/
├── bin/                    # Compiled binary output
│   └── api                 # Built executable
├── cmd/
│   └── api/
│       └── main.go         # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go       # Configuration management
│   ├── db/
│   │   └── db.go           # Database connection and management
│   └── handlers/
│       ├── healthz.go      # Health check endpoint handler
│       └── home.go         # Home endpoint handler
├── .env                    # Environment variables
├── .gitignore              # Git ignore patterns
├── go.mod                  # Go module definition
├── go.sum                  # Go dependency checksums
├── makefile                # Build automation
└── README.md               # Project documentation
```

---

## API Endpoints

### Base URL

```
https://olx-api-wzl0.onrender.com
```

### Endpoints

| Method | Endpoint | Description | Response Format |
|--------|----------|-------------|-----------------|
| GET | `/` | Welcome message for the OLX API | JSON |
| GET | `/healthz` | Health check endpoint to verify API is running | JSON |

### Detailed API Documentation

#### 1. Home Endpoint

**Request:**
```http
GET /
```

**Response:**
```json
{
  "message": "Welcome to the OLX API"
}
```

**Status Code:** `200 OK`

**Headers:**
- `Content-Type: application/json`

---

#### 2. Health Check Endpoint

**Request:**
```http
GET /healthz
```

**Response:**
```json
{
  "status": "ok"
}
```

**Status Code:** `200 OK`

**Headers:**
- `Content-Type: application/json`

**Usage:**
- Monitoring application health
- Load balancer health checks
- Container orchestration (Kubernetes, Docker) readiness probes
- CI/CD pipeline status verification

---

## Setup Guide

### Prerequisites

Before you begin, ensure you have the following installed:

- **Go** (version 1.26.7 or higher)
  - Download: [https://go.dev/dl/](https://go.dev/dl/)
  - Verify: `go version`

- **Git** (for version control)
  - Download: [https://git-scm.com/downloads](https://git-scm.com/downloads)

- **Make** (optional, for build automation)

### Installation

1. **Clone the repository** (if not already cloned):
   ```bash
   git clone https://github.com/dineshEth/olx-api.git
   cd olx-api
   ```

2. **Install Go dependencies**:
   ```bash
   go mod download
   ```

3. **Install development tools** (optional):
   ```bash
   go install github.com/joho/godotenv@latest
   ```

### Environment Configuration

Create or modify the `.env` file in the project root:

```bash
# .env
PORT=8080
ENV=development
DATABASE_URL=postgres://username:password@localhost:5432/dbname?sslmode=disable
```

**Environment Variables:**

| Variable | Description | Required | Default |
|----------|-------------|----------|---------|
| `PORT` | The port number the server will listen on | Yes | - |
| `ENV` | Environment mode (development, staging, production) | Yes | - |
| `DATABASE_URL` | PostgreSQL connection string | Yes | - |

**For Production:**
```bash
PORT=8080
ENV=production
DATABASE_URL=postgres://user:password@host:5432/olx?sslmode=require
```

**For Testing:**
```bash
PORT=3000
ENV=testing
DATABASE_URL=postgres://user:password@localhost:5432/olx_test?sslmode=disable
```

### Running the Application

#### Option 1: Using Makefile (Recommended)

```bash
# Build and run the application
make run

# Or build only
make build

# Then run the binary
./bin/api
```

#### Option 2: Direct Go Execution

```bash
# Build the application
go build -o ./bin/api ./cmd/api/

# Run the binary
./bin/api
```

#### Option 3: Development Mode

```bash
# Run directly with Go (auto-rebuilds on changes)
go run ./cmd/api/
```

**Server Output:**
```
Server starting on port 8080...
```

---

## Usage

Once the server is running, you can access the API endpoints:

### Using curl

```bash
# Test the home endpoint
curl http://localhost:8080/

# Check health status
curl http://localhost:8080/healthz
```

### Using Browser

Open your browser and navigate to:
- [http://localhost:8080/](http://localhost:8080/)
- [http://localhost:8080/healthz](http://localhost:8080/healthz)

### Testing with Live Deployment

You can also test the live deployment:

```bash
# Test live health check
curl https://olx-api-wzl0.onrender.com/healthz

# Test live home endpoint
curl https://olx-api-wzl0.onrender.com/
```

---

## Project Structure

### Architecture Overview

```
┌─────────────────────────────────────┐
│              Client                  │
│   (Browser, Postman, cURL, etc.)    │
└─────────────────────────────────────┘
                    │
                    ▼
┌─────────────────────────────────────┐
│           HTTP Server                │
│         (net/http)                   │
│  ┌─────────────┐  ┌─────────────┐   │
│  │   /         │  │  /healthz   │   │
│  │ HomeHandler │  │ HealthzHandler│  │
│  └─────────────┘  └─────────────┘   │
└─────────────────────────────────────┘
                    │
                    ▼
┌─────────────────────────────────────┐
│            Configuration             │
│        (internal/config)              │
│  - Loads environment variables       │
│  - Manages application settings      │
└─────────────────────────────────────┘
                    │
                    ▼
┌─────────────────────────────────────┐
│            Database                  │
│          (PostgreSQL)                │
│  - Connection pooling                │
│  - Persistent data storage           │
└─────────────────────────────────────┘
```

### Code Organization

- **cmd/api/main.go**: Main application entry point with server configuration and database connection
- **internal/config/config.go**: Environment configuration and validation
- **internal/db/db.go**: Database connection management and PostgreSQL integration
- **internal/handlers/*.go**: HTTP request handlers for all endpoints

### Design Patterns

- **Clean Architecture**: Separation of concerns with clear boundaries
- **Dependency Injection**: Configuration passed to handlers as needed
- **Standard Library First**: Minimal external dependencies

---

## Development

### Adding New Endpoints

1. Create a new handler file in `internal/handlers/`:
   ```go
   package handlers
   
   import "net/http"
   
   func NewEndpointHandler(w http.ResponseWriter, r *http.Request) {
       w.Header().Set("Content-Type", "application/json")
       w.WriteHeader(http.StatusOK)
       w.Write([]byte(`{"message": "New endpoint"}`))
   }
   ```

2. Register the handler in `cmd/api/main.go`:
   ```go
   mux.HandleFunc("GET /new-endpoint", handlers.NewEndpointHandler)
   ```

### Build for Production

```bash
# Build for Linux
GOOS=linux GOARCH=amd64 go build -o ./bin/api ./cmd/api/

# Build for Windows
GOOS=windows GOARCH=amd64 go build -o ./bin/api.exe ./cmd/api/
```

---

## Deployment

The application is currently deployed on **Render** at:
[https://olx-api-wzl0.onrender.com](https://olx-api-wzl0.onrender.com)

For local deployment or other platforms:

1. **Docker** (coming soon)
2. **Kubernetes** (coming soon)
3. **Systemd Service** (for production servers)

---

## Contributing

Contributions are welcome! Please follow these guidelines:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

---

## License

This project is private and proprietary. All rights reserved.

---

## Support

For questions, issues, or support, please contact:

**Developer**: Dinesh Prajapati

**Email**: dineshprajapati@example.com (replace with actual email)

---

## Version History

| Version | Date | Description |
|---------|------|-------------|
| 0.2.0 | November 2024 | Added PostgreSQL database integration with connection pooling |
| 0.1.0 | November 2024 | Initial release with health check and home endpoints |

---

*Built with ❤️ using Go and the OLX platform*
