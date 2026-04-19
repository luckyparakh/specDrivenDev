# Requirements — Phase 2: Database & Migrations

## Goal

Introduce persistent storage. A SQLite database file is created on startup, schema migrations are version-controlled, and type-safe query stubs exist for the `agents` table.

---

## Scope

### In scope

- Add `modernc.org/sqlite` (pure-Go SQLite driver, no CGO)
- Add `golang-migrate/migrate` CLI and library
- Create `db/migrations/` directory holding versioned SQL migration files
- Initial migration: `agents` table with columns `id`, `name`, `email`, `password_hash`, `created_at`
- Add `make migrate` target that runs all pending migrations against `db.sqlite`
- Configure `sqlc` and generate query stubs for the `agents` table into `db/`

### Out of scope

- Sessions table (Phase 3)
- Auth handlers (Phase 3)
- Any UI changes
- Seeding data

---

## Key Decisions

| Decision | Choice | Rationale |
|---|---|---|
| SQLite driver | `modernc.org/sqlite` | Pure Go, no CGO — single-binary constraint |
| DB file location | `db.sqlite` at project root | Simple for MVP; easy to relocate later |
| Migrations directory | `db/migrations/` | Co-located with sqlc output; one `db/` package owns all data-layer artifacts |
| sqlc output package | `db/` (`package db`) | Single place for generated query code; keeps `main.go` clean |
| Migration tool | `golang-migrate` | SQL-native, version-controlled files, simple `make` integration |
| sqlc version | Latest stable (v1) | Sufficient for the query patterns in Phase 2–5 |

---

## Context

AgentClinic is a satirical AI wellness portal (see `specs/mission.md`). The stack targets a **single deployable binary** with no external runtime dependencies beyond a SQLite file (see `specs/tech-stack.md`). Phase 2 lays the data foundation that Phases 3–7 will build on.

The `agents` table stores one row per registered AI agent. `password_hash` is stored instead of a plaintext password; bcrypt hashing is implemented in Phase 3.

---

## Constraints

- No CGO (driver must be `modernc.org/sqlite`)
- Migration files must be plain SQL (no Go-level migration logic)
- sqlc-generated files are committed to the repo; not gitignored
- `make migrate` must be idempotent (re-runs are safe)
