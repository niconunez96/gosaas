package main

import (
	"context"
	authHttp "_compiled/modules/auth/infra/http"
	authPages "_compiled/modules/auth/web/pages"
	"_compiled/modules/shared/infra/htmx"
	sharedHttp "_compiled/modules/shared/infra/http"
	"_compiled/modules/shared/infra/telemetry"
	sharedPages "_compiled/modules/shared/web/pages"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()
	_ = godotenv.Load()
	t, err := telemetry.SetupTelemetry(ctx)
	if err != nil {
		log.Println("Could not load telemetry")
	}
	defer func() {
		if err := t.Shutdown(context.Background()); err != nil {
			log.Fatalf("failed to shutdown TracerProvider: %v", err)
		}
	}()

	globalRouter := chi.NewRouter()

	globalRouter.Use(htmx.HTMXRequestMiddleware())
	globalRouter.Use(authHttp.UnauthorizedMiddleware(authPages.UnauthorizedPage(), authPages.UnauthorizedPage()))

	globalRouter.Use(t.LoggerMiddleware())
	globalRouter.Use(t.TracerMiddleware())
	globalRouter.Use(chimiddleware.Recoverer)

	// Load static files
	globalRouter.Handle(
		"/public/*",
		sharedHttp.DisableCacheMiddleware(http.StripPrefix("/public/", http.FileServer(http.FS(os.DirFS("public"))))),
	)

	globalRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		htmx.Render(r, w, sharedPages.LandingPage())
	})
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		globalRouter.ServeHTTP(w, r)
	})
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	telemetry.LogDebug(fmt.Sprintf("Starting server on port %s", port))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), handler); err != nil {
		telemetry.LogError(fmt.Sprintf("Failed to start server on port %s", port), err)
		os.Exit(1)
	}
}
