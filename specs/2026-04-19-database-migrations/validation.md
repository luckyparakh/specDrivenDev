# Validation — Phase 2: Database & Migrations

Phase 2 is done and mergeable when **all three checkpoints below pass**.

---

## Checkpoint 1 — Migration Runs Clean and Schema Is Correct

**How to verify:**

```sh
make migrate
sqlite3 db.sqlite ".schema agents"
```

**Expected output from `.schema`:**

```sql
CREATE TABLE agents (
    id            INTEGER PRIMARY KEY AUTOINCREMENT,
    name          TEXT    NOT NULL,
    email         TEXT    NOT NULL UNIQUE,
    password_hash TEXT    NOT NULL,
    created_at    DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
```

**Pass criteria:**
- `make migrate` exits 0 with no errors
- `db.sqlite` is created at the project root
- The `agents` table exists with the exact five columns above
- Re-running `make migrate` is idempotent (exits 0, no duplicate-table error)

---

## Checkpoint 2 — Go Test: DB Opens and Agents Table Is Queryable

**How to verify:**

```sh
go test ./db/...
```

**Pass criteria:**
- All tests in `db/db_test.go` pass
- Test opens the database (or an in-memory `:memory:` DB with migrations applied)
- Test asserts the `agents` table exists in `sqlite_master`
- Test asserts all five expected columns are present (`id`, `name`, `email`, `password_hash`, `created_at`)

---

## Checkpoint 3 — sqlc Generated Files Compile

**How to verify:**

```sh
go build ./...
go vet ./...
```

**Pass criteria:**
- `go build ./...` exits 0
- `go vet ./...` exits 0
- All sqlc-generated files (`db/db.go`, `db/models.go`, `db/agents.sql.go`, `db/querier.go`) are present and committed
- No CGO in the build graph (verify with `CGO_ENABLED=0 go build ./...`)

---

## Definition of Done

| Check | Command | Must Pass |
|---|---|---|
| Migration runs clean | `make migrate` | Yes |
| Schema matches spec | `sqlite3 db.sqlite ".schema agents"` | Yes |
| DB test passes | `go test ./db/...` | Yes |
| Full build passes | `go build ./...` | Yes |
| Vet passes | `go vet ./...` | Yes |
| No CGO | `CGO_ENABLED=0 go build ./...` | Yes |
| `db.sqlite` gitignored | `git status db.sqlite` shows ignored | Yes |
| sqlc files committed | `git status db/` shows no untracked | Yes |
