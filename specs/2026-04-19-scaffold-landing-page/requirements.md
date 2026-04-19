# Requirements — Phase 1: Project Scaffold & Landing Page

## Scope

This phase delivers a working local server that renders a branded, mobile-responsive landing page. No database, no auth, no dynamic data. Everything is static HTML produced by Templ templates styled with Tailwind CSS and templui components.

Delivery is in two steps: a **minimal home page** (Group 4) that proves the full pipeline end-to-end, followed by the **full landing page** (Group 5) that adds all required sections.

---

## Decisions

### Module & Repository
- Go module: `github.com/luckyparakh/agentclinic`
- Go version: **1.24.10**
- Repository root is also the Go module root

### HTTP Server
- Framework: **Gin** (`github.com/gin-gonic/gin`)
- Port: `8080` (configurable via `PORT` env var, defaulting to `8080`)
- Static assets served from `static/` at the `/static/` URL prefix

### Templating & Styling
- Templates: **Templ** — all HTML is written in `.templ` files, never raw Go string literals
- Components: **templui** — use its Button and Card components for the feature strip and CTA; do not hand-roll equivalents
- CSS: **Tailwind CSS v4** — compiled to `static/css/app.css`; no inline styles, no `<style>` blocks
- JS: **htmx** loaded from CDN in `base.templ`; no other JavaScript

### Branding (locked)
| Element | Value |
|---|---|
| Site name | AgentClinic |
| Tagline | *"Relief for the Overworked AI"* |
| Primary color | Deep teal (`teal-700` / `teal-800`) |
| Background | Off-white (`neutral-50`) |
| Accent / CTA | Warning amber (`amber-400` / `amber-500`) |
| Body font | System font stack (no custom font loading in Phase 1) |

### Minimal Home Page (required — Group 4)

The `GET /` route must render a styled Templ page containing:
- Site name ("AgentClinic") as a teal-colored heading
- Tagline (*"Relief for the Overworked AI"*) as a subtitle
- No nav, no cards, no footer — only enough to prove Go → Gin → Templ → Tailwind → browser works end-to-end

### Full Landing Page Sections (required — Group 5)
1. **Nav** — site name left-aligned; "Book a Session" CTA right-aligned (links to `#` for now)
2. **Hero** — full-width section, large h1 with site name, tagline as subtitle, one amber CTA button
3. **Feature strip** — 3 equal cards: "Describe Your Ailment", "Choose a Therapy", "Book Your Session"; each with a short satirical one-liner
4. **Footer** — minimal: copyright year, site name, one satirical aside

### Tooling
- `make dev` runs **both** `templ generate --watch` and `air` as concurrent processes in a single target; a developer should only need one terminal command
- `.air.toml` watches `.go` and `.templ` output files; excludes `tmp/`, `bin/`, `static/`
- `npm run css` drives Tailwind; managed via a minimal `package.json` (no framework, no bundler)

---

## Out of Scope for This Phase
- Database, migrations, or any persistent storage
- Authentication or sessions
- Any dynamic route beyond `GET /` and `GET /health`
- Custom fonts or icon libraries
- Dark mode
- Email or notifications

---

## Context
- Mission: playful/satirical tone — copy should reflect that from day one
- Audience includes course students and conference demo developers; the scaffold must be readable and easy to fork
- Refer to [tech-stack.md](../tech-stack.md) for authoritative tool choices and [mission.md](../mission.md) for tone guidance
