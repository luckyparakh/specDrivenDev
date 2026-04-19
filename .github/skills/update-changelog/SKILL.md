---
name: update-changelog
description: "Update CHANGELOG.md before merging. Use when: updating changelog, writing changelog entry, recording changes before merge, logging what changed today, summarising work done. Reads git log and open diffs to produce dated bullet points, then writes them to CHANGELOG.md."
argument-hint: "Optional: scope or phase label (e.g. 'Phase 1', 'auth')"
---

# Update Changelog

Produces a new dated section in `CHANGELOG.md` summarising all changes since the last entry.
Invoke this skill manually before merging a branch.

---

## When to Use

- You are about to merge and want to record what changed
- You want to add a dated entry for today's work
- `CHANGELOG.md` doesn't exist yet (bootstrap from git history)

---

## Procedure

### Step 1 — Read the current state

1. Read `CHANGELOG.md` to find the most recent dated heading.
2. Run the git log script to list commits since that date:
   ```
   git log --pretty=format:"%ad %s" --date=short
   ```
3. Run the diff script to see uncommitted or unpushed changes:
   ```
   git diff HEAD
   git status --short
   ```

### Step 2 — Draft the entry

Produce a new `## YYYY-MM-DD` heading for today's date.

Write **one bullet per logical change**, not one per commit. Group related commits into a single bullet. Use the imperative mood ("Add X", "Fix Y", "Update Z").

Rules:
- Lead with the user-visible effect, not the implementation detail
- Keep each bullet to one line
- If the argument hint was provided (e.g. "Phase 1"), prefix relevant bullets with that label
- Do not include merge commits, version bumps, or changelog-update commits themselves

### Step 3 — Insert into CHANGELOG.md

Insert the new section **immediately after the header block** (the lines before the first `---`), above any existing dated sections. Do not remove existing entries.

Format:
```markdown
## YYYY-MM-DD

- Added X
- Fixed Y
- Updated Z

---
```

### Step 4 — Confirm

Show the user the new section and ask for approval before writing if the diff is ambiguous. If the content is clear, write directly.

---

## Bootstrap (no CHANGELOG.md exists)

If `CHANGELOG.md` is absent:
1. Run `git log --pretty=format:"%ad %s" --date=short` to get all commits
2. Group them by date
3. Synthesise one bullet per logical group per date
4. Create `CHANGELOG.md` with the standard header, then all dated sections oldest-to-newest at the bottom, newest at the top
5. Write the file

---

## CHANGELOG.md Header Template

```markdown
# Changelog

All notable changes to AgentClinic are recorded here.
Entries are added manually before each merge using the `update-changelog` skill.

---
```

---

## Rules

| Rule | Detail |
|---|---|
| Date format | `## YYYY-MM-DD` |
| Order | Newest date at the top |
| Granularity | One bullet per logical change, not per commit |
| Tone | Imperative, concise, user-facing |
| Never include | Merge commits, changelog-update-only commits, CI noise |
