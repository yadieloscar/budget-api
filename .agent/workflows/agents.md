---
description: Guidelines for AI agents working on the budget-api Go project
---

# Agent Guidelines for Budget API

## Project Overview

This is a Go-based REST API for budget management, following a layered architecture.

## Directory Structure

- `cmd/api/main.go` - Application entry point
- `internal/handlers/` - HTTP handlers and routing (Gin)
- `internal/services/` - Business logic layer
- `internal/repo/` - Database persistence (DAO layer)
- `internal/models/` - Domain models and DTOs
- `internal/middleware/` - HTTP middleware
- `configs/` - Viper configuration helpers
- `db_docker/` - Docker Compose for local Postgres
- `docs/` - Project documentation

## Key Conventions

1. **Layered Architecture**:
   - `handlers/` - HTTP adapters, route registration
   - `services/` - Business logic
   - `repo/` - Database persistence
   - `models/` - Domain types

2. **API Versioning**: Routes mount under `/api/v1`

3. **Database**: PostgreSQL is required

## Common Tasks

### Adding a New Feature

1. Add model types to `internal/models/`
2. Add repository interface and implementation to `internal/repo/`
3. Add service interface and implementation to `internal/services/`
4. Add handler and route registration to `internal/handlers/`
5. Wire up in `internal/handlers/routes.go`

### Running the Application

// turbo
```bash
go run cmd/api/main.go
```

### Running Tests

// turbo
```bash
go test ./...
```

### Starting Local Database

// turbo
```bash
docker-compose -f db_docker/docker-compose.yml up -d
```

## Code Style

- Use standard Go formatting (`go fmt`)
- Follow Go naming conventions (camelCase for private, PascalCase for public)
- Use meaningful error messages
- Add comments for exported functions
