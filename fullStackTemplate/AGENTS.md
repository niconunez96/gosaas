# Full Stack Template Agent Gateway

Entry point for AI agents working at template root.

## Scope Routing

- `front/` only -> follow `front/AGENTS.md`
- `back/` only -> follow `back/AGENTS.md`
- both -> treat as integration work

Precedence:
1. User instruction
2. Nearest scoped guide
3. This gateway

## Repository Map

- `front/` React + TypeScript frontend
- `back/` Go backend
- `Justfile` common tasks

## Integration Contract

1. Confirm API contract first.
2. Implement backend behavior.
3. Integrate frontend with documented response shapes.
4. Validate loading, error, and empty states.

## Safety Rules

- Keep changes minimal and scoped.
- Avoid product-specific content in shared template assets.
- If ownership is unclear, ask before implementing.
