# Plan — Phase 1: Project Scaffold & Landing Page

Each task group is independently completable and results in a passing `go build`.

---

## Group 1 — Go Module & Bare Server

1. `go mod init github.com/luckyparakh/agentclinic`
2. Create `main.go`: start Gin on `:8080`, single `GET /health` route returning `200 OK`
3. Confirm `go build ./...` succeeds and `./agentclinic` starts

---

## Group 2 — Makefile & Live Reload

1. Add `Makefile` with targets:
   - `make build` — runs `templ generate` then `go build -o bin/agentclinic .`
   - `make dev` — runs `templ generate --watch` and `air` concurrently (use `make -j2` or a shell background job)
   - `make lint` — `go vet ./...`
2. Add `.air.toml` (default config, watch `.go` and `.templ` files, exclude `tmp/`)
3. Confirm `make dev` starts without error and restarts on a `.go` file save

---

## Group 3 — Templ + templui + Tailwind Pipeline

1. `go get github.com/a-h/templ`
2. Add `templui` dependency (`go get github.com/axzilla/templui`)
3. Add `tailwind.config.js` (content globs: `**/*.templ`, `**/*.go`)
4. Add `package.json` with `tailwindcss` dev dependency; add `npm run css` script
5. Extend `make dev` to also run `npm run css -- --watch`
6. Create `static/css/app.css` (Tailwind directives: `@tailwind base/components/utilities`)
7. Create `web/layouts/base.templ` — HTML shell with `<link>` to compiled CSS and htmx CDN script tag
8. Confirm Tailwind output file is generated and a background color class applies visibly

---

## Group 4 — Minimal Home Page

1. Create `web/pages/home.templ` — extends `base.templ`; renders only:
   - Site name heading ("AgentClinic") using a teal text class
   - Tagline (*"Relief for the Overworked AI"*) as a subtitle paragraph
2. Wire `GET /` Gin route to render `home.templ`
3. Serve `static/` directory via Gin's `Static` middleware
4. Confirm `http://localhost:8080` returns a styled page with the site name and tagline — nothing else required at this step

---

## Group 5 — Full Landing Page

1. Expand `web/pages/home.templ` (already wired from Group 4) — add all required sections:
   - **Nav**: site name "AgentClinic" (left), placeholder "Book a Session" CTA link (right)
   - **Hero**: large heading, tagline (*"Relief for the Overworked AI"*), primary CTA button
   - **Features strip**: 3 cards — "Describe Your Ailment", "Choose a Therapy", "Book Your Session"
   - **Footer**: copyright line with a satirical note
2. Apply templui component variants for button and card where available

---

## Group 6 — Mobile Responsiveness

1. Audit all layout containers: ensure `max-w-*`, `px-4`, responsive grid classes (`sm:`, `md:`) are in place
2. Verify nav collapses cleanly at < 640 px (hamburger or stacked layout)
3. Verify feature cards stack to single column on mobile

---

## Group 7 — Test Suite

1. Refactor `main.go`: extract `setupRouter() *gin.Engine` so tests can wire the router without starting a real server
2. Write `main_test.go` (package `main`) — table-driven integration tests:
   - `GET /health` → `200 OK`
   - `GET /` → `200 OK` + all required text present in the HTML body
   - Unknown routes → `404`
3. Export `FeatureCard` struct and `featureCards` slice from `web/pages/home.templ` (capitalise field names)
4. Write `web/pages/home_test.go` (package `pages`) — table-driven component unit tests:
   - `TestFeatureCardsData` — asserts count, titles, and taglines of the static card data
   - `TestFeatureCardView` — renders each card variant and asserts content in the HTML buffer
   - `TestHomeComponent` — renders the full `Home()` component and asserts all sections are present
5. Confirm `go test ./...` exits `0` with no failures
