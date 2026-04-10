import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { render, screen } from '@testing-library/react'
import { App } from './App'

const renderApp = () => {
  const queryClient = new QueryClient({
    defaultOptions: {
      queries: {
        retry: false,
      },
    },
  })

  return render(
    <QueryClientProvider client={queryClient}>
      <App />
    </QueryClientProvider>,
  )
}

describe('App', () => {
  it('renders the default starter content', () => {
    renderApp()

    expect(
      screen.getByRole('heading', { name: /application ready\./i }),
    ).toBeInTheDocument()
    expect(
      screen.getByText(/neutral starter view/i),
    ).toBeInTheDocument()
    expect(
      screen.getByRole('button', { name: /continue/i }),
    ).toBeInTheDocument()
  })
})
