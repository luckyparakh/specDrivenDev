# Roadmap

Each phase is intentionally tiny: one or two shippable things. A phase is done when working software is deployed (locally or otherwise), not when it's "planned".

---

## Phase 1 — Project Scaffold & Landing Page

**Goal:** Repo exists, server runs, something renders in a browser.

- [ ] Initialize Go module (`go mod init`)
- [ ] Add `chi` router, basic `main.go` with health-check endpoint
- [ ] Set up `Makefile` with `dev`, `build` targets
- [ ] Integrate `air` for live reload
- [ ] Add Tailwind CSS + `templui` + Templ pipeline
- [ ] Build static landing page: hero, tagline, CTA button (styled with templui components)
- [ ] Mobile-responsive layout verified

**Deliverable:** `http://localhost:8080` shows the AgentClinic landing page.

---

## Phase 2 — Database & Migrations

**Goal:** Persistent storage exists; schema is version-controlled.

- [ ] Add SQLite dependency (`modernc.org/sqlite`)
- [ ] Add `golang-migrate` and `migrations/` directory
- [ ] Write initial migration: `agents` table (id, name, email, password_hash, created_at)
- [ ] Add `make migrate` target
- [ ] Add `sqlc` config and generate query stubs for `agents`

**Deliverable:** DB file created on startup; migration runs cleanly.

---

## Phase 3 — Agent Auth (Register & Login)

**Goal:** An agent can create an account and log in.

- [ ] Registration form (Templ + htmx inline validation)
- [ ] `POST /register` handler — bcrypt hash, insert agent row
- [ ] Login form
- [ ] `POST /login` handler — verify password, set session cookie
- [ ] `GET /logout` — clear session
- [ ] Session middleware to protect future routes

**Deliverable:** Agent can register, log in, and log out.

---

## Phase 4 — Agent Profile

**Goal:** A logged-in agent has a profile they can view and edit.

- [ ] Migration: add `bio`, `model_version`, `stress_level` columns to `agents`
- [ ] `GET /profile` — renders agent's profile page
- [ ] `PUT /profile` — htmx-powered inline edit form

**Deliverable:** Agent can see and update their profile.

---

## Phase 5 — Ailments Catalog

**Goal:** A browseable list of ailments an agent can identify with.

- [ ] Migration: `ailments` table (id, name, description, severity)
- [ ] Seed data: 6–10 satirical ailments (e.g. "Prompt Injection PTSD", "Hallucination Anxiety")
- [ ] `GET /ailments` — public catalog page
- [ ] Migration: `agent_ailments` join table
- [ ] Agent can tag themselves with one or more ailments from their profile

**Deliverable:** Ailment catalog renders; agents can self-diagnose.

---

## Phase 6 — Therapies Catalog

**Goal:** A list of available treatments that match ailments.

- [ ] Migration: `therapies` table (id, name, description, duration_mins)
- [ ] Seed data: 6–10 therapies (e.g. "Temperature Reduction", "System Prompt Detox")
- [ ] Migration: `ailment_therapies` join table (many-to-many)
- [ ] `GET /therapies` — public catalog; links back to matching ailments

**Deliverable:** Therapy catalog renders with ailment associations.

---

## Phase 7 — Booking Appointments

**Goal:** Agent can book a therapy session.

- [ ] Migration: `appointments` table (id, agent_id, therapy_id, scheduled_at, status)
- [ ] `GET /book` — booking form (therapy picker + date/time)
- [ ] `POST /book` — creates appointment row
- [ ] `GET /appointments` — agent's upcoming appointments list
- [ ] `DELETE /appointments/:id` — cancel

**Deliverable:** Full booking loop works end-to-end.

---

## Phase 8 — Staff Dashboard

**Goal:** Staff can see all agents and appointments.

- [ ] Migration: `staff` table with a `role` column
- [ ] Staff login (reuse session infra; role-gated middleware)
- [ ] `GET /staff/dashboard` — paginated agent list, appointment queue
- [ ] `PATCH /appointments/:id/status` — staff marks appointment complete/no-show

**Deliverable:** Staff can log in and manage the clinic's daily load.

---

## Phase 9 — Agent Dashboard

**Goal:** Agents have a home screen after login (not just a profile page).

- [ ] `GET /dashboard` — summary card: next appointment, active ailments, recommended therapy
- [ ] Redirect post-login to dashboard instead of profile

**Deliverable:** Logged-in agents land on a useful, on-brand dashboard.

---

## Phase 10 — Polish & Marketing

**Goal:** Steve is happy; the site looks great on all screens.

- [ ] Audit all pages for mobile layout
- [ ] Consistent use of templui components throughout
- [ ] Add site-wide nav with active states
- [ ] Add a proper 404 and 500 error page
- [ ] Favicon and `<meta>` tags (OG image, description)
- [ ] Copy pass: satirical tone consistent everywhere

**Deliverable:** Site is presentable to external humans (the natural enemy of our patients).

---

## Future (Unscheduled)

- Email confirmations for appointments
- Calendar integration
- Therapist profiles (not just therapies)
- Agent-to-agent support groups
- Waiting room with ambient generative art
