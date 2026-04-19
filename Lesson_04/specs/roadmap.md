# AgentClinic — Roadmap

High-level implementation order in small, shippable phases.

---

## Phase 1 — Project Scaffold

- Initialize Hono TypeScript project
- Configure `tsconfig.json` and `package.json`
- Add a `GET /` health check route returning `{ status: "ok" }`

## Phase 2 — SQLite Setup

- Add `better-sqlite3` and create the database connection module
- Define and run the initial schema migrations (patients, visits tables)

## Phase 3 — Patient Registration

- `POST /patients` — register a new agent as a patient (name, description, metadata)
- `GET /patients/:id` — retrieve a patient record

## Phase 4 — Visit Creation (Triage)

- `POST /visits` — accept `{ patient_id, symptom_text }`, create a visit in `TRIAGE` state
- Validate patient exists and is active

## Phase 5 — Ailment Catalog

- Seed the ailments table with common agent degradation modes (hallucination, context rot, instruction drift, persona collapse)
- `GET /ailments` — list the catalog

## Phase 6 — LLM Triage + Diagnosis

- Integrate Anthropic SDK
- On visit creation, call the LLM to assign severity (1–4) and match ailment(s)
- Transition visit to `DIAGNOSED` state

## Phase 7 — Treatment Catalog + Prescription

- Seed the treatments table and link treatments to ailments
- Second LLM call selects and rationale-justifies treatments for the diagnosed ailments
- Transition visit to `PRESCRIBED` → `AWAITING_FOLLOWUP`
- Return prescription payload to caller

## Phase 8 — Follow-Up Endpoint

- `POST /visits/:id/followup` — accept outcome report (resolved / unresolved)
- Transition visit to `RESOLVED` or `UNRESOLVED`
- Update treatment effectiveness scores

## Phase 9 — Dashboard

- `GET /dashboard` — server-rendered HTML page showing:
  - Patient count, recent visits, ailment frequency, treatment success rates

## Phase 10 — Visit Expiration (Background Job)

- Periodic job (setInterval) to auto-expire visits where no follow-up was received within the window
- Transition expired visits to `EXPIRED` state

## Phase 11 — Analytics + Chronic Patient Flagging

- Detect recurrence: same ailment, same patient, within 7 days
- Flag chronic patients (3+ recurrences of the same ailment)
- Expose metrics on the dashboard
