package gotuna

import (
	"net/http"
	"time"
)

// Logging middleware is used to log every requests to the app's Logger.
func (app App) Logging() MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			start := time.Now()

			next.ServeHTTP(w, r)

			if app.Logger != nil {
				app.Logger.Debugf("%s %s finished in %s", r.Method, r.URL.Path, time.Since(start))
			}
		})
	}
}
