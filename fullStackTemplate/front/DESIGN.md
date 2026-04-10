# Frontend Design Strategy (Template)

## Goal

Provide a clean default visual baseline that teams can customize quickly without
rewriting foundation components.

## Principles

1. Keep surfaces simple and readable.
2. Use semantic tokens (`primary`, `surface`, `ink`) instead of brand names.
3. Prefer composition from shared primitives over screen-specific styling.
4. Optimize for clarity before visual flair.

## Token Model

- Color tokens should express intent, not brand (`primary`, `error`, `info`).
- Radius, spacing, and shadows should be centralized in `src/styles.css`.
- Reusable classes should use `foundation-*` naming.

## Layout Baseline

- Mobile-first.
- Content should remain legible across breakpoints.
- Shared shells (`Topbar`, `SidebarNav`, `Card`) should be reused before
  introducing custom layout primitives.

## Component Baseline

Use the primitives in `src/shared/components/ui` as the default layer:

- `Button`, `Badge`, `Card`
- `Input`, `Textarea`, `Toggle`, `Slider`
- `Toast`, `DataTable`, `EmptyState`, `Topbar`, `SidebarNav`

## Customization Rule

When adapting this template for a new project:

1. Update tokens first.
2. Re-skin shared primitives second.
3. Build feature-specific views last.

Avoid hardcoding project branding directly in foundational components.
