# Plan — Phase 2: Database & Migrations

Numbered task groups in order. Complete each group before starting the next.

---

## Group 1 — Add SQLite Dependency

1. Run `go get modernc.org/sqlite` to add the pure-Go SQLite driver.
2. Run `go mod tidy` and verify `go.mod` / `go.sum` are updated.
3. Confirm no CGO references appear in the build output (`go build ./...`).

---

## Group 2 — Set Up golang-migrate

1. Add `github.com/golang-migrate/migrate/v4` to `go.mod` (`go get github.com/golang-migrate/migrate/v4`).
2. Add the SQLite source driver: `go get github.com/golang-migrate/migrate/v4/database/sqlite`.
3. Create `db/migrations/` directory.
4. Add a `migrate` target to `Makefile`:
   ```makefile
   migrate:
       migrate -source file://db/migrations -database sqlite3://db.sqlite up
   ```
5. Verify `make migrate` runs without error (no migrations yet — that's fine at this step).

---

## Group 3 — Write Initial Migration

1. Create `db/migrations/000001_create_agents.up.sql`:
   ```sql
   CREATE TABLE IF NOT EXISTS agents (
       id            INTEGER PRIMARY KEY AUTOINCREMENT,
       name          TEXT    NOT NULL,
       email         TEXT    NOT NULL UNIQUE,
       password_hash TEXT    NOT NULL,
       created_at    DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
   );
   ```
2. Create `db/migrations/000001_create_agents.down.sql`:
   ```sql
   DROP TABLE IF EXISTS agents;
   ```
3. Run `make migrate` — `db.sqlite` should be created and the `agents` table present.
4. Verify with `sqlite3 db.sqlite ".schema agents"` (or equivalent Go test in Group 5).

---

## Group 4 — Configure sqlc

1. Create `db/sqlc.yaml`:
   ```yaml
   version: "2"
   sql:
     - engine: "sqlite"
       queries: "db/queries/"
       schema: "db/migrations/"
       gen:
         go:
           package: "db"
           out: "db"
   ```
2. Create `db/queries/agents.sql` with initial query stubs:
   ```sql
   -- name: CreateAgent :one
   INSERT INTO agents (name, email, password_hash)
   VALUES (?, ?, ?)
   RETURNING *;

   -- name: GetAgentByEmail :one
   SELECT * FROM agents WHERE email = ? LIMIT 1;

   -- name: GetAgentByID :one
   SELECT * FROM agents WHERE id = ? LIMIT 1;
   ```
3. Run `sqlc generate` from the project root.
4. Confirm generated files appear in `db/` (`db.go`, `models.go`, `agents.sql.go`, `querier.go`).

---

## Group 5 — Go Test: DB Opens and Schema Is Correct

1. Create `db/db_test.go` with a table-driven test that:
   - Opens `db.sqlite` (or an in-memory `:memory:` DB with migrations applied)
   - Queries `sqlite_master` to assert the `agents` table exists
   - Asserts all expected columns (`id`, `name`, `email`, `password_hash`, `created_at`) are present
2. Run `go test ./db/...` — all tests must pass.

---

## Group 6 — Build & Housekeeping

1. Run `go build ./...` — generated sqlc files must compile without errors.
2. Add `db.sqlite` to `.gitignore` (DB file must not be committed).
3. Ensure `db/migrations/`, `db/queries/`, and all sqlc-generated Go files *are* committed.
4. Run `go vet ./...` — no issues.
