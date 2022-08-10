package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"

	"github.com/Slava12/gotuna"
	"github.com/Slava12/gotuna/examples/fullapp"
	"github.com/Slava12/gotuna/examples/fullapp/i18n"
	"github.com/Slava12/gotuna/examples/fullapp/static"
	"github.com/Slava12/gotuna/examples/fullapp/views"
	"github.com/Slava12/gotuna/test/doubles"
	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
)

func main() {

	port := ":8888"
	keyPairs := os.Getenv("APP_KEY")
	cookieStore := sessions.NewCookieStore([]byte(keyPairs))
	cookieStore.Options.HttpOnly = true         // more secure
	cookieStore.Options.MaxAge = 30 * 24 * 3600 // expire in one month

	app := fullapp.MakeApp(fullapp.App{
		App: gotuna.App{
			Router:         gotuna.NewMuxRouter(),
			Logger:         nil,
			UserRepository: doubles.NewUserRepositoryStub(),
			Session:        gotuna.NewSession(cookieStore, "app_session"),
			Static:         static.EmbededStatic,
			StaticPrefix:   "",
			ViewFiles:      views.EmbededViews,
			Locale:         gotuna.NewLocale(i18n.Translations),
		},
		Csrf: csrf.Protect(
			[]byte(keyPairs),
			csrf.FieldName("csrf_token"),
			csrf.CookieName("csrf_token"),
		),
	})

	fmt.Printf("starting server at http://localhost%s \n", port)

	if err := http.ListenAndServe(port, app.Router); err != nil {
		log.Fatalf("could not listen on port %s %v", port, err)
	}
}
