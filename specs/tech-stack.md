# Tech Stack

## Guiding Principle

Prefer boring, proven technology. The stack should be deployable as a **single binary** with no external runtime dependencies beyond the database file.

---

## Server

| Concern | Choice | Rationale |
|---|---|---|
| Language | **Go 1.24.10** | Fast, simple, strong stdlib, single binary output |
| HTTP router | **Gin** (`github.com/gin-gonic/gin`) | Fast, widely known, great for demos and teaching; familiar to most Go developers |
| HTML templating | **Templ** | Type-safe Go templates, compile-time checked |
| UI components | **templui** | Pre-built Templ component library; pairs naturally with Tailwind |
| Partial updates | **htmx** | Server-driven interactivity without a JS framework |

## Styling

| Concern | Choice | Rationale |
|---|---|---|
| CSS utility framework | **Tailwind CSS v4** | Co-located styles, tree-shaken output, mobile-first |
| Component source | **templui** | Ships Tailwind-compatible components, no custom CSS required for basics |

## Data

| Concern | Choice | Rationale |
|---|---|---|
| Database | **SQLite** (via `modernc.org/sqlite` — pure Go, no CGO) | Zero-ops, single file, perfect for MVP; easy to swap to Postgres later |
| Migrations | **`golang-migrate`** | SQL migration files, version-controlled, simple CLI |
| Query layer | **`sqlc`** | Generates type-safe Go code from SQL queries; no ORM magic |

## Auth

| Concern | Choice |
|---|---|
| Strategy | Cookie-based sessions (HTTP-only, SameSite=Strict) |
| Password hashing | `bcrypt` (via `golang.org/x/crypto`) |
| Session store | SQLite-backed session table |

> No OAuth or third-party SSO in the initial phases. Agents register with a made-up email and a password they immediately forget.

## Build & Tooling

| Tool | Purpose |
|---|---|
| `templ generate` | Compiles `.templ` files to Go |
| `tailwindcss` CLI | Builds and purges CSS |
| `air` | Live-reload during development |
| `make` | Single `Makefile` with `dev`, `build`, `migrate` targets |

## Deployment

- Single Go binary + static assets (embedded via `embed.FS`)
- Docker-friendy: `scratch` base image, <20 MB image target
- SQLite database file mounted as a volume
- No external services required at launch

---

## Constraints & Non-Goals (Phase 1–3)

- No JavaScript framework (React, Vue, etc.) — htmx only
- No cloud-specific services (S3, RDS, etc.)
- No email sending — stubs only until a later phase
- No payments
