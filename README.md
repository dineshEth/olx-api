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
  - [Running Database Migrations](#running-database-migrations)
- [Database Schema](#database-schema)
- [Usage](#usage)
- [Project Structure](#project-structure)

---

## Features

- **Health Check Endpoint**: Monitor application status and availability
- **Home Endpoint**: Welcome message and API information
- **Listings Endpoint**: Retrieve classified advertisements from the database
- **Environment Configuration**: Secure environment variable management
- **Database Connectivity**: PostgreSQL database integration with connection pooling
- **Database Migrations**: Version-controlled schema management with up/down migrations
- **Production-Ready**: Configured with proper timeouts for production deployment
- **Lightweight**: Minimal dependencies for fast startup and low memory footprint

---

## Tech Stack

| Category | Technology | Version |
|----------|------------|---------|
| **Language** | Go (Golang) | 1.26.7 |
| **Framework** | Standard Library (net/http) | - |
| **Database** | PostgreSQL | - |
| **Database Driver** | pgx v5 | v5.13.0 |
| **Migration Tool** | golang-migrate | v4.20.1 |
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
│   ├── api/
│   │   └── main.go         # Application entry point
│   └── migrate/
│       └── main.go         # Database migration CLI tool
├── internal/
│   ├── config/
│   │   └── config.go       # Configuration management
│   ├── db/
│   │   └── db.go           # Database connection and management
│   └── handlers/
│       ├── healthz.go      # Health check endpoint handler
│       └── home.go         # Home endpoint handler
├── migrations/             # Database migration files
│   ├── 000001_create_listing_table.up.sql   # Create listings table
│   └── 000001_create_listing_table.down.sql # Drop listings table
├── .env                    # Environment variables
├── .gitignore              # Git ignore patterns
├── go.mod                  # Go module definition
├── go.sum                  # Go dependency checksums
├── makefile                # Build automation with migration commands
└── README.md               # Project documentation
```

---

## API Endpoints

### Base URL

```
https://olx-api-wzl0.onrender.com
```

### Endpoints

All endpoints return JSON responses.

| Method | Endpoint | Description | Response Format |
|--------|----------|-------------|-----------------|
| GET | `/` | Welcome message for the OLX API | JSON |
| GET | `/healthz` | Health check endpoint to verify API is running | JSON |
| GET | `/listings` | Retrieve all classified listings from the database | JSON |

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

#### 3. Listings Endpoint

**Request:**
```http
GET /listings
```

**Response:**
```json
[
  {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "title": "Used Car for Sale",
    "description": "2018 model, excellent condition",
    "price": "15000.00",
    "city": "New York",
    "created_at": "2024-11-13T10:30:00Z"
  },
  {
    "id": "550e8400-e29b-41d4-a716-446655440001",
    "title": "Laptop",
    "description": "MacBook Pro 2020",
    "price": "1200.00",
    "city": "San Francisco",
    "created_at": "2024-11-12T14:20:00Z"
  }
]
```

**Status Code:** `200 OK`

**Headers:**
- `Content-Type: application/json`

**Query Parameters:**
- None currently supported (returns up to 20 most recent listings)

**Features:**
- Returns all listings from the database
- Ordered by creation date (newest first)
- Limited to 20 results for performance
- Includes full listing details: id, title, description, price, city, and creation timestamp

---

## Setup Guide

### Prerequisites

Before you begin, ensure you have the following installed:

- **Go** (version 1.26.7 or higher)
  - Download: [https://go.dev/dl/](https://go.dev/dl/)
  - Verify: `go version`

- **Git** (for version control)
  - Download: [https://git-scm.com/downloads](https://git-scm.com/downloads)

- **PostgreSQL** (database server)
  - Download: [https://www.postgresql.org/download/](https://www.postgresql.org/download/)
  - Verify: `psql --version`

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

4. **Install database migration tool** (required for migrations):
   ```bash
   go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
   ```

### Environment Configuration

Create or modify the `.env` file in the project root:

```bash
# .env
PORT=8080
ENV=development
DATABASE_URL=postgres://username:password@localhost:5432/olx?sslmode=disable
```

**Environment Variables:**

| Variable | Description | Required | Default |
|----------|-------------|----------|---------|
| `PORT` | The port number the server will listen on | Yes | - |
| `ENV` | Environment mode (development, staging, production) | Yes | - |
| `DATABASE_URL` | PostgreSQL connection string (pgx format) | Yes | - |

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

**Note:** The DATABASE_URL must use the pgx connection string format. Example: `postgres://user:password@localhost:5432/database?sslmode=disable`

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

### Running Database Migrations

#### Using Makefile

```bash
# Apply all pending migrations (up)
make migrate-up

# Rollback the last migration (down)
make migrate-down
```

#### Using Migration CLI Directly

```bash
# Apply all pending migrations
go run ./cmd/migrate up

# Rollback the last migration
go run ./cmd/migrate down
```

**Migration Commands:**
- `up`: Apply all pending migrations to bring the database schema up to date
- `down`: Rollback the most recent migration (one step back)

**Important:** Always run migrations before starting the application to ensure the database schema is current.

**Server Output:**
```
Server starting on port 8080...
```

---

## Database Schema

The application uses PostgreSQL with the following schema:

### Tables

#### listings

The primary table for storing classified advertisements.

| Column | Type | Description | Constraint |
|--------|------|-------------|------------|
| `id` | UUID | Unique identifier | PRIMARY KEY, DEFAULT gen_random_uuid() |
| `title` | TEXT | Listing title | NOT NULL |
| `description` | TEXT | Detailed description | NOT NULL |
| `price` | NUMERIC(10, 2) | Price of the item | NOT NULL |
| `city` | TEXT | Location city | NOT NULL |
| `created_at` | TIMESTAMPTZ | Creation timestamp | NOT NULL, DEFAULT NOW() |

**Index:**
- Primary key on `id` (UUID)

### Creating New Migrations

To add a new migration:

1. Create a new migration file in the `migrations/` directory following the naming convention:
   - `YYYYMMDD_HHMMSS_description.up.sql` - Migration up script
   - `YYYYMMDD_HHMMSS_description.down.sql` - Migration down script

2. Write your SQL in the up file to apply changes
3. Write the reverse SQL in the down file to rollback changes

4. Run the migration:
   ```bash
   go run ./cmd/migrate up
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

# Get all listings
curl http://localhost:8080/listings
```

### Using Browser

Open your browser and navigate to:
- [http://localhost:8080/](http://localhost:8080/)
- [http://localhost:8080/healthz](http://localhost:8080/healthz)
- [http://localhost:8080/listings](http://localhost:8080/listings)

### Testing with Live Deployment

You can also test the live deployment:

```bash
# Test live health check
curl https://olx-api-wzl0.onrender.com/healthz

# Test live home endpoint
curl https://olx-api-wzl0.onrender.com/

# Get live listings
curl https://olx-api-wzl0.onrender.com/listings
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
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐   │
│  │   /         │  │  /healthz   │  │  /listings  │   │
│  │ HomeHandler │  │ HealthzHandler│  │ ListingsHandler│  │
│  └─────────────┘  └─────────────┘  └─────────────┘   │
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
│  - Connection pooling (pgx)          │
│  - Persistent data storage           │
│  - listings table                    │
└─────────────────────────────────────┘
                    │
                    ▼
┌─────────────────────────────────────┐
│         Migrations                   │
│  (golang-migrate/v4)                 │
│  - Version-controlled schema         │
│  - Up/Down migration support         │
│  - File-based migration storage      │
└─────────────────────────────────────┘
```

### Code Organization

- **cmd/api/main.go**: Main application entry point with server configuration, database connection, and handler registration
- **cmd/migrate/main.go**: Database migration CLI tool using golang-migrate
- **internal/config/config.go**: Environment configuration and validation (includes DATABASE_URL)
- **internal/db/db.go**: Database connection management using pgx driver with connection pooling
- **internal/handlers/home.go**: Home endpoint handler
- **internal/handlers/healthz.go**: Health check endpoint handler
- **internal/handlers/listings.go**: Listings endpoint handler with database queries
- **migrations/*.sql**: Version-controlled database schema migrations (up/down scripts)

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

### Adding New Database Tables

1. Create new migration files in `migrations/`:
   ```bash
   # Up migration (create table)
   migrations/YYYYMMDD_HHMMSS_create_table_name.up.sql
   
   # Down migration (drop table)
   migrations/YYYYMMDD_HHMMSS_create_table_name.down.sql
   ```

2. Run the migration:
   ```bash
   go run ./cmd/migrate up
   ```

3. Use the table in your application through the database connection

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
| 0.4.0 | November 2024 | Added /listings endpoint to retrieve classified advertisements from database |
| 0.3.0 | November 2024 | Added database migrations with golang-migrate, pgx driver, and listings table schema |
| 0.2.0 | November 2024 | Added PostgreSQL database integration with connection pooling |
| 0.1.0 | November 2024 | Initial release with health check and home endpoints |

---

*Built with ❤️ using Go and the OLX platform*
