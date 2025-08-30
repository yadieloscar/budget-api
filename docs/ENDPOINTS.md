# Endpoints

Base path: `/api/v1`

## Budget

### GET `/budget/`

Returns the current user's budget.

- Auth: requires `user_id` present in Gin context (middleware TBD).
- Response: `200 OK` with `Budget` JSON, `404 Not Found` if none, `401 Unauthorized` if `user_id` missing.

Example:

```
curl http://localhost:8080/api/v1/budget/ \
  -H 'X-Debug-User: USER-123'    # example only; add real middleware
```

Response (200):

```
{
  "id": "...",
  "user_id": "USER-123",
  "name": "September Budget",
  "month": "2024-09",
  "status": "active",
  "available_amount": 500000,
  "total_amount": 500000,
  "currency": "USD",
  "created_at": "2024-08-21T17:10:00Z",
  "updated_at": "2024-08-21T17:10:00Z"
}
```

### POST `/budget/create`

Creates a budget for the current user.

- Auth: requires `user_id` present in Gin context (middleware TBD).
- Body:

```
{
  "user_id": "USER-123",
  "name": "September Budget",
  "total_amount": 500000,
  "currency": "USD",
  "month": "2024-09"
}
```

- Responses:
  - `201 Created` with created `Budget` JSON
  - `400 Bad Request` if validation fails (e.g., non‑positive amount)
  - `401 Unauthorized` if `user_id` missing
  - `500 Internal Server Error` on unexpected errors

Notes:

- The repository writes `status = "active"` on creation and sets `available_cents = total_amount_cents`.
- Database requires a `budgets` table; see README for the DDL sketch.

