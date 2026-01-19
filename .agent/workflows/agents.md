---
description: Guidelines for AI agents working on the budget-api Go project
---

# Agent Guidelines for Budget API

## Project Overview

This is a Go-based REST API for budget management, following a feature-module architecture.

## Directory Structure

- `cmd/api/main.go` - Application entry point
- `internal/api/` - API layer with router and middleware
- `internal/api/<feature>/` - Feature modules (handler, service, repo, model, module)
- `configs/` - Viper configuration helpers
- `db_docker/` - Docker Compose for local Postgres
- `docs/` - Project documentation

## Key Conventions

1. **Feature Module Pattern**: Each feature has:
   - `handler/` - HTTP adapters (Gin handlers)
   - `service/` - Business logic
   - `repo/` - Database persistence
   - `model/` - Domain models and DTOs
   - `module/` - Dependency wiring

2. **API Versioning**: Routes mount under `/api/v1`

3. **Database**: PostgreSQL is required

## Common Tasks

### Adding a New Feature Module

1. Create directory `internal/api/<feature>/`
2. Add `model/` with request DTOs and domain models
3. Add `repo/` with SQL persistence layer
4. Add `service/` with business logic
5. Add `handler/` with Gin HTTP handlers
6. Add `module/` to wire dependencies and register routes
7. Register the module in the main router

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