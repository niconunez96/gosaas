package authHttp

import (
	"net/http"

	"_compiled/modules/shared/infra/htmx"
	sharedHttp "_compiled/modules/shared/infra/http"

	"github.com/a-h/templ"
)

func UnauthorizedMiddleware(
	unauthorizedPage templ.Component,
	forbiddenPage templ.Component,
) sharedHttp.Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			responseWriter := sharedHttp.NewStatusCodeRecorder(w)
			h.ServeHTTP(responseWriter, r)
			switch responseWriter.StatusCode {
			case http.StatusUnauthorized:
				htmx.Render(r, w, unauthorizedPage)
				return
			case http.StatusForbidden:
				htmx.Render(r, w, forbiddenPage)
				return
			}
		})
	}
}
