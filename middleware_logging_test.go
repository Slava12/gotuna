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

func TestLogging(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/sample", nil)
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

	logger := app.Logging()
	logger(http.NotFoundHandler()).ServeHTTP(response, request)

	assert.Contains(t, wlog.String(), "GET")
	assert.Contains(t, wlog.String(), "/sample")
}
