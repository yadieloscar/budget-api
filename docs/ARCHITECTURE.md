# Architecture

This service follows a feature‑module structure with clear separation of concerns.

## Layers

- Entry (`cmd/api/main.go`): loads env, opens DB, builds router, starts server.
- API (`internal/api`): top‑level router and middleware.
  - Shared middleware (`internal/api/shared/middleware`): common HTTP middleware.
  - Versioning: routes are mounted under `/api/v1`.
- Feature modules (example: `internal/api/budget`):
  - `handler`: HTTP adapters (Gin) that validate/bind requests and shape responses.
  - `service`: business logic and validation; orchestrates repositories.
  - `repo`: persistence; executes SQL using `database/sql`.
  - `model`: domain models and request DTOs.
  - `module`: wires repo → service → handler and registers routes.

Other directories:

- `configs`: Viper‑based configuration helpers (not currently used by main).
- `internal/models`: legacy/shared models not tied to a specific feature.
- `pkg/utils`: placeholder for shared utilities.
- `db_docker`: Docker Compose for local Postgres.

## Request Flow (Budget)

1. Client hits `POST /api/v1/budget/create`.
2. Gin routes to `handler.BudgetHandler.CreateBudget`.
3. Handler binds JSON to `model.CreateBudgetRequest` and reads `user_id` from context.
4. `service.BudgetService.CreateBudget` validates business invariants.
5. `repo.BudgetRepository.CreateBudget` persists to Postgres and returns the entity.
6. Handler returns the created budget JSON.

## Authentication

Currently missing. Handlers expect `user_id` in the Gin context; add auth middleware to populate it.

## Database

Postgres is expected. The `budgets` table is required by the budget repository.

## Versioning Strategy

All routes mount under `/api/v1`. When breaking changes are introduced, add `/api/v2` alongside `/api/v1` during migration.

