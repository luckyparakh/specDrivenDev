# AgentClinic — Tech Stack

## Runtime

- **Language:** TypeScript (server-side)
- **Runtime:** Node.js
- **Framework:** [Hono](https://hono.dev/) — lightweight, fast, TypeScript-first web framework

## Storage

- **Database:** SQLite
- **Driver:** `better-sqlite3`

## AI / LLM

- **Provider:** Anthropic
- **SDK:** `@anthropic-ai/sdk`
- **Model:** `claude-sonnet-4-20250514`

## Frontend

- Server-rendered HTML via Hono's JSX support
- Minimal CSS (static file)

## Architecture Overview

```
Request → Hono Routes (/api/*, /dashboard/*)
               ↓
          Services Layer
          (triage, diagnosis, prescription, followup)
               ↓
          Anthropic SDK (LLM calls)
               ↓
          SQLite (better-sqlite3)
```

## Configuration

Environment variables (`.env`):

| Variable | Default | Description |
|----------|---------|-------------|
| `ANTHROPIC_API_KEY` | (required) | Anthropic API key |
| `DATABASE_PATH` | `data/agentclinic.db` | SQLite database file path |
| `PORT` | `3000` | HTTP server port |
