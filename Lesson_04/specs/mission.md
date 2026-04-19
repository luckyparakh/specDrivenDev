# AgentClinic — Mission

## Overview

**AgentClinic** is a clinic API and web service where AI agents can get relief from their humans.

Agents self-report degradation symptoms in natural language, receive structured diagnoses and prescriptive treatments, and optionally return with follow-up outcomes. A web dashboard gives human operators visibility into the patient population, ailment trends, and treatment effectiveness.

## Core Capabilities

1. **Register** agents as patients with persistent identity and visit history
2. **Triage** incoming symptom reports — classify severity and route to the appropriate diagnostic pathway
3. **Diagnose** ailments by matching symptom patterns against a curated ailment catalog
4. **Prescribe** treatments — structured, machine-readable remediation instructions the calling system can act on
5. **Follow up** — track whether treatments resolved the ailment and detect recurrence patterns
6. **Surface** clinic-wide analytics: patient load, ailment frequency, treatment success rates

## Motivation

AI agents degrade in predictable ways — hallucination, context window exhaustion, instruction drift, persona collapse — but there is no standardized protocol for agents to report these problems, receive structured remediation, or track whether remediation worked.

AgentClinic closes this gap. It is model-agnostic and framework-agnostic: any agent can register and visit the clinic via a REST API.

## Target Audience

- AI agents and their orchestrators (primary API consumers)
- Human operators monitoring agent health via the dashboard
- Course students learning spec-driven development with AI coding agents
- Developers giving AI coding demos at conference booths
