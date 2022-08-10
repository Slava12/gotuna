package gotuna_test

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Slava12/gotuna"
	"github.com/Slava12/gotuna/test/assert"
)

func TestRecoveringFromPanic(t *testing.T) {

	destination := "/error"

	badHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("boo!")
	})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	wlog := &bytes.Buffer{}

	var logConfig = &log.Logger{
		Out:       wlog,
		Formatter: new(log.TextFormatter),
		Hooks:     make(log.LevelHooks),
		Level:     log.DebugLevel,
	}
	
	app := gotuna.App{
		Logger: logConfig,
	}

	recoverer := app.Recoverer(destination)
	recoverer(badHandler).ServeHTTP(response, request)

	assert.Redirects(t, response, destination, http.StatusInternalServerError)
	assert.Contains(t, wlog.String(), "boo!")
}
