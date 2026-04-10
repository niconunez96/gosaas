# Frontend Template Agent Guide

This file defines the default frontend baseline for `front/`.

## Mission

- Keep architecture simple and understandable.
- Preserve reusable primitives and state boundaries.
- Avoid product-specific branding in template code.

## Approved Stack

- Vite
- React + TypeScript
- Tailwind CSS
- React Router
- TanStack Query (server state)
- Zustand (client/UI state)
- React Hook Form + Zod (forms)
- pnpm

## State Boundaries

1. Server state -> TanStack Query
2. UI state -> Zustand
3. Form state -> React Hook Form + Zod

Do not mix responsibilities in one store.

## UI Primitives Policy

Use and extend shared primitives under `src/shared/components/ui` before creating feature-local UI.

## Styling Rules

1. Tailwind-first.
2. Prefer centralized tokens in `src/styles.css`.
3. Keep reusable classes generic; avoid product naming.

## Validation Commands

- Canonical frontend check: `just front-check`
- Under the hood:
  - `pnpm --dir front format:check`
  - `pnpm --dir front lint`
  - `pnpm --dir front test`
