# Go Base API

Go REST API using the [Fiber](https://gofiber.io/) framework running in a Docker container. The API includes API Key authentication and follows clean architecture principles.

## 📋 Table of Contents

- [Features](#features)
- [Architecture](#architecture)
- [Quick Start](#quick-start)
- [API Endpoints](#api-endpoints)
- [Authentication](#authentication)
- [Configuration](#configuration)
- [Docker](#docker)
- [Database](#database)
- [Development](#development)

## ✨ Features

- ✅ **Fiber Framework** - Fast and minimalist Go web framework
- ✅ **API Key Authentication** - Secure authentication via X-API-Key header
- ✅ **Clean Architecture** - Separate layers: routes, controllers, services, repository, models
- ✅ **PostgreSQL Database** - Persistent data storage with Repository Pattern
- ✅ **Docker Support** - Multi-stage Dockerfile and docker-compose
- ✅ **Environment Configuration** - .env file support
- ✅ **Standardized Responses** - Unified JSON response format
- ✅ **Health Check** - API status check endpoint
- ✅ **CRUD Operations** - Sample user management
- ✅ **Query Builder** - Squirrel SQL builder for type-safe queries

## 🏗️ Architecture

```
api/
├── cmd/
│   ├── server/
│   │   └── main.go              # Entry point
│   └── seed/
│       └── main.go              # Seed command
├── internal/
│   ├── config/
│   │   └── config.go            # Configuration loading
│   ├── database/
│   │   └── database.go          # Database connection
│   ├── middleware/
│   │   └── auth.go              # API Key authentication
│   ├── routes/
│   │   └── routes.go            # Route definitions
│   ├── controllers/
│   │   ├── health_controller.go # Health check
│   │   └── user_controller.go   # User CRUD
│   ├── services/
│   │   └── user_service.go      # Business logic
│   ├── repository/
│   │   ├── user_repository.go   # Repository interface
│   │   └── postgres_user_repository.go  # PostgreSQL implementation
│   ├── models/
│   │   └── user.go              # Data models
│   └── utils/
│       ├── response.go          # Response helpers
│       └── validator.go         # Validation
├── migrations/
│   ├── 001_create_users_table.sql
│   ├── 002_create_indexes.sql
│   └── seeds/
│       └── 001_seed_users.sql
├── .env                         # Environment variables
├── .env.example                 # Example .env
├── Dockerfile                   # Docker image
├── docker-compose.yml           # Docker Compose (with PostgreSQL)
├── go.mod                       # Go module
├── Makefile                     # Command line commands
└── README.md                    # Documentation
```

### Architecture Layers

| Layer | Folder | Description |
|-------|--------|-------------|
| **Entry Point** | `cmd/server/` | Server startup |
| **Seed Command** | `cmd/seed/` | Database seeding with sample data |
| **Config** | `internal/config/` | Configuration management |
| **Database** | `internal/database/` | Database connection |
| **Middleware** | `internal/middleware/` | Authentication, logging |
| **Routes** | `internal/routes/` | URL routing |
| **Controllers** | `internal/controllers/` | HTTP request handling |
| **Services** | `internal/services/` | Business logic |
| **Repository** | `internal/repository/` | Database queries (Repository Pattern) |
| **Models** | `internal/models/` | Data structures |
| **Utils** | `internal/utils/` | Helper functions |
| **Migrations** | `migrations/` | SQL files for creating tables and indexes |

## 🚀 Quick Start

### Prerequisites

- [Go 1.21+](https://golang.org/dl/) (for local development)
- [Docker](https://www.docker.com/) (for running in container)
- [Docker Compose](https://docs.docker.com/compose/)

### Local Development

1. **Clone the project**
   ```bash
   cd api
   ```

2. **Copy .env file**
   ```bash
   cp .env.example .env
   ```

3. **Edit .env file** - replace API_KEY with your value

4. **Install dependencies**
   ```bash
   go mod tidy
   ```

5. **Run the server**
   ```bash
   go run cmd/server/main.go
   ```

Server will start at `http://localhost:3000`

### Docker

1. **Run with Docker Compose**
   ```bash
   docker-compose up -d
   ```

2. **Check status**
   ```bash
   docker-compose ps
   ```

3. **View logs**
   ```bash
   docker-compose logs -f api
   ```


## 🔗 API Endpoints

### Health Check

| Method | Path | Authentication | Description |
|--------|------|----------------|-------------|
| GET | `/api/health` | ❌ No | API status check |

**Response:**
```json
{
  "success": true,
  "data": {
    "status": "healthy",
    "version": "1.0.0"
  },
  "message": "API is running"
}
```

### Users

| Method | Path | Authentication | Description |
|--------|------|----------------|-------------|
| GET | `/api/users` | ✅ Yes | Get all users |
| GET | `/api/users/:id` | ✅ Yes | Get single user |
| POST | `/api/users` | ✅ Yes | Create user |
| PUT | `/api/users/:id` | ✅ Yes | Update user |
| DELETE | `/api/users/:id` | ✅ Yes | Delete user |

#### GET /api/users
```bash
curl -H "X-API-Key: your-api-key" http://localhost:3000/api/users
```

#### GET /api/users/:id
```bash
curl -H "X-API-Key: your-api-key" http://localhost:3000/api/users/1
```

#### POST /api/users
```bash
curl -X POST \
  -H "X-API-Key: your-api-key" \
  -H "Content-Type: application/json" \
  -d '{"name": "Alice Johnson", "email": "alice@example.com"}' \
  http://localhost:3000/api/users
```

#### PUT /api/users/:id
```bash
curl -X PUT \
  -H "X-API-Key: your-api-key" \
  -H "Content-Type: application/json" \
  -d '{"name": "Alice Updated"}' \
  http://localhost:3000/api/users/1
```

#### DELETE /api/users/:id
```bash
curl -X DELETE \
  -H "X-API-Key: your-api-key" \
  http://localhost:3000/api/users/1
```

## 🔐 Authentication

The API uses **API Key authentication** via the `X-API-Key` header.

### How it works

1. Every protected endpoint request must include the `X-API-Key` header
2. Server compares the key in the header with the one in configuration
3. If key is missing or incorrect → **401 Unauthorized**
4. If key is correct → request continues to the controller

### Exceptions

- `/api/health` endpoint does not require authentication

### Example

```bash
# Correct - 200 OK
curl -H "X-API-Key: your-secret-api-key" http://localhost:3000/api/users

# Wrong key - 401 Unauthorized
curl -H "X-API-Key: wrong-key" http://localhost:3000/api/users

# Missing key - 401 Unauthorized
curl http://localhost:3000/api/users
```

### .env Example

```env
SERVER_PORT=3000
API_KEY=my-super-secret-api-key-12345
ENVIRONMENT=development

# Database Configuration
DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=go_base_api
DB_SSLMODE=disable
```

## 🐳 Docker

### Dockerfile

Uses **multi-stage build** for optimization:
- **Builder stage**: Go compilation
- **Runtime stage**: Alpine Linux, binary only

### docker-compose.yml

The project includes two services:
- **api**: Go API server (port 3000)
- **postgres**: PostgreSQL database (port 5432)

### Docker Commands

```bash
# Build and run
docker-compose up -d --build

# Build only
docker build -t go-base-api .

# Run single container
docker run -p 3000:3000 -e API_KEY=secret go-base-api
```

#### Connecting to Database

```bash
# Via Docker container
docker exec -it go-base-api-postgres psql -U postgres -d go_base_api

# Or locally (if psql is installed)
psql -h localhost -U postgres -d go_base_api
```

#### Viewing Tables

```sql
-- All tables
\dt

-- Table structure
\d users

-- Table indexes
\di
```


#### Exit

```sql
\q
```

### Migrations and Seeding

```bash
# Run migrations (create tables)
make migrate

# Run seed (add sample data)
make seed

# Or directly
go run cmd/seed/main.go
```

### Database Management with Makefile

```bash
# Start all containers
make up

# Stop containers
make down

# Build and start
make build

# View logs
make logs

# Container status
make ps
```



### Adding a New Endpoint

1. **Create model** (`internal/models/`)
2. **Create service** (`internal/services/`)
3. **Create controller** (`internal/controllers/`)
4. **Register route** (`internal/routes/routes.go`)

## 📝 License

MIT
