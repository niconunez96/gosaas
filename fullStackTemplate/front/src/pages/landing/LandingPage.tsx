import { Button } from '../../shared/components/ui/button'
import { Card } from '../../shared/components/ui/card'
import { Topbar } from '../../shared/components/ui/topbar'

export function LandingPage() {
  return (
    <div className="foundation-page">
      <header className="sticky top-0 z-30 px-4 pt-4 sm:px-6 lg:px-8">
        <Topbar
          actions={
            <Button className="min-h-10 px-5 py-2 text-[11px] uppercase tracking-section">
              Primary Action
            </Button>
          }
          className="mx-auto max-w-6xl rounded-full border border-white/70 bg-white/80 px-5 py-3 shadow-overshoot backdrop-blur-xl lg:px-6"
          eyebrow="Full Stack Template"
          title="Frontend Starter"
        />
      </header>

      <main className="mx-auto flex max-w-4xl flex-col gap-8 px-4 pb-16 pt-10 sm:px-6 lg:px-8 lg:pt-14">
        <Card className="space-y-4 bg-white">
          <p className="foundation-section-eyebrow">Status</p>
          <h1 className="text-4xl font-black tracking-tight text-ink">
            Application ready.
          </h1>
          <p className="foundation-body max-w-2xl">
            This is a neutral starter view to confirm the frontend bootstraps
            correctly.
          </p>
          <div className="pt-2">
            <Button size="lg">Continue</Button>
          </div>
        </Card>
      </main>

      <footer className="mx-auto mt-8 max-w-4xl rounded-t-[2.5rem] bg-surface px-4 py-10 sm:px-6 lg:px-8">
        <p className="max-w-md text-[11px] font-extrabold uppercase tracking-section text-ink/45">
          Full Stack Template • Starter Baseline
        </p>
      </footer>
    </div>
  )
}
