package gotuna

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

func clientIP(r *http.Request) string {
	clientIP := strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if len(clientIP) > 0 {
		return clientIP
	}
	clientIP = r.Header.Get("X-Forwarded-For")
	if index := strings.IndexByte(clientIP, ','); index >= 0 {
		clientIP = clientIP[0:index]
	}
	clientIP = strings.TrimSpace(clientIP)
	if len(clientIP) > 0 {
		return clientIP
	}
	return strings.TrimSpace(r.RemoteAddr)
}

// Logging middleware is used to log every requests to the app's Logger.
func (app App) Logging() MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			start := time.Now()

			next.ServeHTTP(w, r)

			app.Logger = app.Logger.WithFields(log.Fields{
				"client_ip": clientIP(r),
				"method":    r.Method,
				"url":       r.URL.Path,
			})

			if app.Session != nil {
				id, err := app.Session.GetUserID(r)
				if err == nil {
					app.Logger = app.Logger.WithFields(log.Fields{
						"user_id": id,
					})
				}
				edoID, err := app.Session.GetEdoID(r)
				if err == nil {
					app.Logger = app.Logger.WithFields(log.Fields{
						"user_edo_id": edoID,
						"user_name":   app.Session.GetName(r),
					})
				}
			}

			if app.Logger != nil {
				app.Logger.Debugf("finished in %s", time.Since(start))
			}
		})
	}
}
