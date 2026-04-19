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

## Testing

### Philosophy

All tests are written in Go using the standard `testing` package. Tests live **alongside source files** (`_test.go` next to each package). No external test runner.

### Strategy: Table-Driven Tests

Use Go table-driven tests (`[]struct{ name, ... }` + `t.Run`) for every layer:

| Layer | What to test | Tool |
|---|---|---|
| **Business logic** | Pure functions, data transforms | `testing` |
| **HTTP handlers** | Status codes, headers, redirects | `net/http/httptest` + `httptest.NewRecorder` |
| **HTML output** | Rendered Templ components — presence of text, element structure | `net/http/httptest` + stdlib `strings`/`bytes` |
| **Route integration** | Full request → Gin router → rendered HTML | `httptest.NewServer` or `httptest.NewRecorder` with a wired `*gin.Engine` |

### Frontend Testing (Go-only, no browser tooling)

Render Templ components or fire HTTP requests against a test Gin engine, then assert on the HTML string output. No Playwright, no Cypress — the stdlib is sufficient for Phase 1–3.

**Pattern for handler + HTML assertions:**

```go
// Example: table-driven test for GET /
func TestHomeRoute(t *testing.T) {
    router := setupRouter() // returns a configured *gin.Engine

    tests := []struct {
        name         string
        wantStatus   int
        wantContains []string
    }{
        {
            name:       "home page renders landing content",
            wantStatus: http.StatusOK,
            wantContains: []string{
                "AgentClinic",
                "Relief for the Overworked AI",
                "Book a Session",
                "Describe Your Ailment",
            },
        },
    }

    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            w := httptest.NewRecorder()
            req := httptest.NewRequest(http.MethodGet, "/", nil)
            router.ServeHTTP(w, req)

            if w.Code != tc.wantStatus {
                t.Errorf("status: got %d, want %d", w.Code, tc.wantStatus)
            }
            body := w.Body.String()
            for _, s := range tc.wantContains {
                if !strings.Contains(body, s) {
                    t.Errorf("body missing %q", s)
                }
            }
        })
    }
}
```

**Pattern for Templ component unit tests:**

```go
// Example: assert a component renders expected markup
func TestFeatureCardView(t *testing.T) {
    tests := []struct {
        name         string
        card         featureCard
        wantContains []string
    }{
        {
            name: "renders title and body",
            card: featureCard{
                title: "Describe Your Ailment",
                body:  "some body text",
            },
            wantContains: []string{"Describe Your Ailment", "some body text"},
        },
    }

    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            var buf bytes.Buffer
            err := featureCardView(tc.card).Render(context.Background(), &buf)
            if err != nil {
                t.Fatalf("render error: %v", err)
            }
            html := buf.String()
            for _, s := range tc.wantContains {
                if !strings.Contains(html, s) {
                    t.Errorf("rendered HTML missing %q", s)
                }
            }
        })
    }
}
```

> **Rule of thumb:** test behaviour, not markup. Assert on meaningful text content and HTTP semantics — not on specific HTML tag nesting.

## Build & Tooling

| Tool | Purpose |
|---|---|
| `templ generate` | Compiles `.templ` files to Go |
| `tailwindcss` CLI | Builds and purges CSS |
| `air` | Live-reload during development |
| `make` | Single `Makefile` with `dev`, `build`, `migrate` targets |
| `go test ./...` | Runs all table-driven tests |

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
