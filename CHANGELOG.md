# Changelog

All notable changes to AgentClinic are recorded here.
Entries are added manually before each merge using the `update-changelog` skill.

---

## 2026-04-19

- Enhanced responsive design requirements and validation criteria for landing page (viewport meta, breakpoint standard, per-phase mobile checklist in roadmap)
- Initialized project scaffold: Go module, Gin server, Templ + templui + Tailwind CSS v4 pipeline, live-reload with air
- Built minimal home page (Group 4) and full landing page with nav, hero, feature cards, footer (Groups 5-6)
- Added table-driven Go test suite: route integration tests and Templ component unit tests
- Added product-level specifications: mission, tech-stack, roadmap, Phase 1 plan/requirements/validation
- Added testing strategy to tech-stack (Go stdlib, `net/http/httptest`, alongside-source pattern)
- Added `update-changelog` skill and bootstrapped CHANGELOG.md from git history
- Updated prompt.md with Lesson 08 prompts
- Restarted project with clean Go-based stack (replaced earlier TypeScript prototype)
- First commit: initial repository structure
