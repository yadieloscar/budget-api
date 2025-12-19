# Architecture

This service follows a layered architecture with clear separation of concerns.

## Layers

- Entry (`cmd/api/main.go`): loads env, opens DB, builds router, starts server.
- Handlers (`internal/handlers`): HTTP adapters (Gin) that validate/bind requests, shape responses, and define routes.
- Services (`internal/services`): business logic and validation; orchestrates repositories.
- Repositories (`internal/repo`): persistence; executes SQL using `database/sql`.
- Models (`internal/models`): domain models and request DTOs.
- Middleware (`internal/middleware`): common HTTP middleware (CORS, etc.).

Other directories:

- `configs`: Viper‑based configuration helpers (not currently used by main).
- `pkg/utils`: placeholder for shared utilities.
- `db_docker`: Docker Compose for local Postgres.

## Request Flow (Budget)

1. Client hits `POST /api/v1/budget/create`.
2. Gin routes to `handlers.BudgetHandler.CreateBudget`.
3. Handler binds JSON to `models.CreateBudgetRequest` and reads `user_id` from context.
4. `services.BudgetService.CreateBudget` validates business invariants.
5. `repo.BudgetRepository.CreateBudget` persists to Postgres and returns the entity.
6. Handler returns the created budget JSON.

## Authentication

Currently missing. Handlers expect `user_id` in the Gin context; add auth middleware to populate it.

## Database

Postgres is expected. The `budgets` table is required by the budget repository.

## Versioning Strategy

All routes mount under `/api/v1`. When breaking changes are introduced, add `/api/v2` alongside `/api/v1` during migration.
