# Validation — Phase 1: Project Scaffold & Landing Page

Phase 1 is **merge-ready** when every check below passes. Work through them in order; a failure at any step blocks merge.

---

## 1. Build

```
go build ./...
```
- Must exit `0` with no errors or warnings
- Produced binary must be named `bin/agentclinic` when run via `make build`

---

## 2. Dev Toolchain

```
make dev
```
- Must start without error
- Both `templ generate --watch` and `air` must be running as concurrent processes
- Editing and saving a `.go` file must trigger a server restart (visible in terminal output within 3 seconds)
- Tailwind CSS watcher must regenerate `static/css/app.css` on template changes

---

## 3. Server

```
curl -i http://localhost:8080/health
```
- Must return `HTTP/1.1 200 OK`

---

## 4. Minimal Home Page

Open `http://localhost:8080` in a browser.

| Check | Expected |
|---|---|
| Page loads | No 404, no 500, no blank page |
| `<title>` | Contains "AgentClinic" |
| Site name | "AgentClinic" visible as a teal-colored heading |
| Tagline | *"Relief for the Overworked AI"* visible as subtitle text |
| Tailwind applied | Page is not unstyled default browser HTML |

---

## 5. Full Landing Page Content

Open `http://localhost:8080` in a browser after Group 5 is complete.

| Check | Expected |
|---|---|
| Nav | "AgentClinic" visible left; "Book a Session" visible right |
| Hero heading | Site name visible as a prominent `h1` |
| Tagline | *"Relief for the Overworked AI"* visible below the heading |
| CTA button | Amber-colored button present in hero section |
| Feature cards | Exactly 3 cards: "Describe Your Ailment", "Choose a Therapy", "Book Your Session" |
| Footer | Copyright line present with satirical note |

---

## 6. Styling

- All Tailwind classes resolve — no un-styled raw HTML elements
- Primary/accent colors (teal, amber) visible on at least the nav and CTA button
- Background is off-white, not default browser white

---

## 7. Mobile Layout

Resize browser to 375 px width (or use DevTools device emulation).

| Check | Expected |
|---|---|
| Nav | Does not overflow horizontally; items either stack or collapse cleanly |
| Hero text | Readable, not clipped |
| Feature cards | Stack to a single column |
| CTA button | Full-width or appropriately sized; not cut off |

---

## 8. htmx Present

- `htmx` script loaded (visible in browser DevTools > Network, or check page source for `htmx`)
- No JavaScript console errors on page load

---

## 8. `go vet`

```
make lint
```
- Must exit `0` with no reported issues

---

## 9. Test Suite

```
go test ./...
```
- Must exit `0` with no failures
- The following test functions must pass:

| Package | Test | What it covers |
|---|---|---|
| `main` | `TestHealthRoute` | `GET /health` → 200 |
| `main` | `TestHomeRoute` | `GET /` → 200 + all landing page text in body |
| `main` | `TestUnknownRouteReturns404` | undefined routes → 404 |
| `pages` | `TestFeatureCardsData` | static card slice has 3 entries with correct titles and taglines |
| `pages` | `TestFeatureCardView` | each card renders its title, tagline, and body |
| `pages` | `TestHomeComponent` | full `Home()` component renders nav, hero, all 3 cards, footer, htmx script |

---

## Definition of Done

All 8 checks above pass on a clean `git clone` + `make dev` on a machine with Go 1.24.10 and Node.js installed. No manual steps beyond those two commands should be required.
