# Budget API

REST API for personal budgeting, built with Go and Gin. The service is organized by feature modules (e.g., budget) with clear separation between handler, service, and repository layers.

Note: This is a work in progress. Some modules are scaffolds and certain middleware (e.g., authentication) is not implemented yet.

## Project Structure

- `cmd/api`: Application entrypoint (`main.go`).
- `internal/api`: HTTP router, versioning, shared middleware, and feature modules.
  - `internal/api/shared/middleware`: Cross-cutting HTTP middleware (CORS).
  - `internal/api/budget`: Budget feature (handler, service, repo, model, module).
- `internal/models`: Legacy or shared models (kept for reference).
- `configs`: Viper-based configuration helpers (currently unused by main).
- `pkg/utils`: Place for shared helpers to avoid circular dependencies.
- `db_docker`: Docker Compose for local Postgres.

See `docs/ARCHITECTURE.md` for more detail.

## Prerequisites

- Go (1.21+ recommended)
- Postgres (local or Docker)

## Local Database (Docker)

Spin up a Postgres instance with Docker Compose:

```
cd db_docker
docker compose up -d
```

This starts Postgres on `localhost:5432` with database/user `budget` and password `secret`.

## Configuration

The API reads `DB_DSN` from the environment. Example:

```
export DB_DSN='postgres://budget:secret@localhost:5432/budget?sslmode=disable'
```

For local development, a `.env` file is supported (see `.env` in repo):

```
DB_DSN=postgres://budget:secret@localhost:5432/budget?sslmode=disable
```

## Run

```
go mod tidy
go run cmd/api/main.go
```

The server listens on `:8080` by default.

## Current Endpoints

Base path: `/api/v1`

- `GET /budget/` — Get the current user's budget.
- `POST /budget/create` — Create a budget for the current user.

Important: Handlers expect `user_id` to be present in the Gin context. There is currently no middleware setting this value, so these endpoints will return 401 Unauthorized unless you inject `user_id` (e.g., via temporary middleware during local testing).

Minimal request example (replace with your `user_id` injection strategy):

```
curl -X POST http://localhost:8080/api/v1/budget/create \
  -H 'Content-Type: application/json' \
  -d '{
    "user_id": "USER-123",  
    "name": "September Budget",
    "total_amount": 500000,
    "currency": "USD",
    "month": "2024-09"
  }'
```

## Database Schema Expectations

The repository layer assumes a `budgets` table similar to:

```
CREATE TABLE budgets (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id TEXT NOT NULL,
  name TEXT NOT NULL,
  month TEXT NOT NULL,
  status TEXT NOT NULL,
  available_cents BIGINT NOT NULL,
  total_amount_cents BIGINT NOT NULL,
  currency TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

Migrations are not included yet.

## Documentation

- `docs/ARCHITECTURE.md` — High-level architecture and module layout.
- `docs/ENDPOINTS.md` — Endpoint details and request/response examples.

## Contributing

Contributions are welcome! Please open an issue or submit a PR. For significant changes, start with an issue to discuss the approach.
