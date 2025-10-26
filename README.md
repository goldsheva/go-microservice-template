# go-microservice-template

A minimal Go microservice skeleton. It includes a Gin-based HTTP server, a healthcheck endpoint, configuration loading from `config.env`, structured logging with Logrus, and a GORM-based data layer (MySQL/PostgreSQL/SQLite).

## Quick start

1) Clone

```bash
git clone https://github.com/goldsheva/go-microservice-template.git
cd go-microservice-template
```

2) Modules setup
- Using this repository as-is (keep module path):
```bash
go mod tidy
```
- Creating a new project from this template (set your own module path):
```bash
rm -f go.mod go.sum
go mod init <your/module>
go mod tidy
```

3) Configuration
Create `config.env` (loaded on startup):
```env
LOG_LEVEL=info                 # debug|info|warn|error|fatal|panic
HTTP_PORT=8080
DB_DRIVER=postgres             # mysql|postgres|sqlite3
DB_HOST=localhost
DB_PORT=5432
DB_NAME=app
DB_USER=app
DB_PASSWORD=secret
```

## Build

Build binaries for Linux and macOS:
```bash
make build
```
Outputs:
- `bin/app-linux`
- `bin/app-macos`

Alternative (local only):
```bash
go build -o bin/app cmd/main.go
```

## Dependencies (core)
- Gin (HTTP framework): `github.com/gin-gonic/gin`
- GORM (ORM): `gorm.io/gorm`
- DB drivers: `gorm.io/driver/mysql`, `gorm.io/driver/postgres`, `gorm.io/driver/sqlite`
- Logrus (logging): `github.com/sirupsen/logrus`
- GORM + Logrus hook: `github.com/onrik/gorm-logrus`
- Godotenv (env loader): `github.com/joho/godotenv`
- Ozzo Validation (config validation): `github.com/go-ozzo/ozzo-validation/v4`

See `go.mod` for the full list.

## HTTP healthcheck
- Route: `GET /api/healthcheck`
- Port: taken from `HTTP_PORT` in `config.env`
- Response: `200 OK`, body: `{ "status": "ok" }`

Example:
```bash
curl -s http://localhost:8080/api/healthcheck | jq .
```

## Useful commands
- Run without building: `go run cmd/main.go`

## Structure
- `cmd/main.go` — application entrypoint
- `internal/workers/http_server.go` — Gin HTTP server and healthcheck route
- `internal/configs/env.go` — environment configuration loader/validation
- `internal/database/*` — GORM initialization (MySQL/PostgreSQL/SQLite)