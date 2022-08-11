package gotuna

import (
	log "github.com/sirupsen/logrus"
	"io/fs"

	"github.com/gorilla/mux"
)

// App is the main app dependency store.
// This is where all the app's dependencies are configured.
type App struct {
	Logger          *log.Entry
	Router          *mux.Router
	Static          fs.FS
	StaticPrefix    string
	ViewFiles       fs.FS
	ViewHelpers     []ViewHelperFunc
	Session         *Session
	EnvironmentName string
	UserRepository  UserRepository
	Locale          Locale
}

// NewMuxRouter returns the underlying mux router instance
func NewMuxRouter() *mux.Router {
	return mux.NewRouter()
}
