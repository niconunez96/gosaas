package htmx

import (
	"context"
	"net/http"

	sharedHttp "_compiled/modules/shared/infra/http"
	sharedPages "_compiled/modules/shared/web/pages"

	"github.com/a-h/templ"
)

type htmxContextKey string

const useFullLayoutPageContextKey htmxContextKey = "use_full_layout_page"

func HTMXRequestMiddleware() sharedHttp.Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("hx-request") != "true" {
				ctx := context.WithValue(
					r.Context(),
					useFullLayoutPageContextKey,
					true,
				)
				r = r.WithContext(ctx)
			}
			h.ServeHTTP(w, r)
		})
	}
}

func shouldRenderFullPage(ctx context.Context) bool {
	useFullLayoutPage, ok := ctx.Value(useFullLayoutPageContextKey).(bool)
	if !ok {
		return false
	}
	return useFullLayoutPage
}

func Render(r *http.Request, w http.ResponseWriter, template templ.Component) {
	ctx := r.Context()
	if shouldRenderFullPage(ctx) {
		ctx := templ.WithChildren(ctx, template)
		_ = sharedPages.LandingPage().Render(ctx, w)
		return
	}
	_ = template.Render(ctx, w)
}
