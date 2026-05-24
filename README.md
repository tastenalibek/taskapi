# taskapi

A RESTful task management API written in Go, backed by PostgreSQL.

## Endpoints

| Method | Path          | Description       |
|--------|---------------|-------------------|
| GET    | /tasks        | List all tasks    |
| POST   | /tasks        | Create a task     |
| GET    | /tasks/{id}   | Get a single task |
| PUT    | /tasks/{id}   | Update a task     |
| DELETE | /tasks/{id}   | Delete a task     |

## Quick start

Requires [Docker](https://docs.docker.com/get-docker/) (make sure Docker Desktop is running).

```bash
git clone https://github.com/tastenalibek/taskapi
cd taskapi
docker compose up --build
```

The API will be available at `http://localhost:8080`.

## Try it out

```bash
# Create a task
curl -s -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"Buy groceries","description":"Milk, eggs, bread"}' | jq

# List all tasks
curl -s http://localhost:8080/tasks | jq

# Update a task
curl -s -X PUT http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Buy groceries","status":"done"}' | jq

# Delete a task
curl -s -X DELETE http://localhost:8080/tasks/1
```

### Task schema

```json
{
  "id": 1,
  "title": "Buy groceries",
  "description": "Milk, eggs, bread",
  "status": "todo",
  "created_at": "2026-05-24T10:00:00Z",
  "updated_at": "2026-05-24T10:00:00Z"
}
```

`status` must be one of: `todo`, `in_progress`, `done`.

## Run locally (without Docker)

```bash
# Start only the database
docker compose up -d db

# Copy and edit env
cp .env.example .env

# Run the server
DATABASE_URL=postgres://taskapi:secret@localhost:5432/taskapi?sslmode=disable go run .
```

## Project structure

```
taskapi/
├── main.go                  # server setup, routing, startup
├── Dockerfile               # multi-stage build (builder + alpine)
├── docker-compose.yml       # postgres + api services
└── internal/
    ├── db/
    │   └── db.go            # connection pool + schema migration
    ├── handler/
    │   └── task.go          # HTTP handlers (CRUD)
    └── model/
        └── task.go          # Task struct + request validation
```

## Tech

- **[chi](https://github.com/go-chi/chi)** — lightweight, idiomatic Go HTTP router
- **`database/sql` + [lib/pq](https://github.com/lib/pq)** — PostgreSQL driver
- **`log/slog`** — structured logging (stdlib, Go 1.21+)
- **Docker Compose** — one-command local environment

## License

MIT
