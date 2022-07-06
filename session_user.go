package gotuna

import (
	"errors"
	"net/http"
)

const (
	// UserIDKey is used as session key to store the current's user unique ID.
	UserIDKey = "_user_id"
	// UserLocaleKey is used as session key for the current's user locale settings.
	UserLocaleKey = "_user_locale"
	// UserEdoIDKey is used as session key to store the current's user EDO ID.
	UserEdoIDKey = "_user_edo_id"
	// UserNameKey is used as session key to store the current's user name.
	UserNameKey = "_user_name"
	// UserEnvironmentKey is used as session key to store the current's user environment.
	UserEnvironmentKey = "_user_environment"
)

// IsGuest checks if current user is not logged in into the app.
func (s Session) IsGuest(r *http.Request) bool {
	id, err := s.GetUserID(r)
	if err != nil || id == "" {
		return true
	}
	return false
}

// SetUserID will save the current user's unique ID into the session.
func (s Session) SetUserID(w http.ResponseWriter, r *http.Request, id string) error {
	return s.Put(w, r, UserIDKey, id)
}

// GetUserID retrieves the current user's unique ID from the session.
func (s Session) GetUserID(r *http.Request) (string, error) {
	id, err := s.Get(r, UserIDKey)
	if err != nil || id == "" {
		return "", errors.New("no user in the session")
	}

	return id, nil
}

// SetLocale will store the user's locale string into the session.
func (s Session) SetLocale(w http.ResponseWriter, r *http.Request, id string) error {
	return s.Put(w, r, UserLocaleKey, id)
}

// GetLocale retrieves the current user's locale string from the session.
func (s Session) GetLocale(r *http.Request) string {
	locale, err := s.Get(r, UserLocaleKey)
	if err != nil {
		return "" // TODO: default locale
	}

	return locale
}

// SetEdoID will store the user's EDO ID string into the session.
func (s Session) SetEdoID(w http.ResponseWriter, r *http.Request, id string) error {
	return s.Put(w, r, UserEdoIDKey, id)
}

// GetEdoID retrieves the current user's EDO ID string from the session.
func (s Session) GetEdoID(r *http.Request) (string, error) {
	edoID, err := s.Get(r, UserEdoIDKey)
	if err != nil {
		return "", errors.New("no EDO ID in the session")
	}

	return edoID, nil
}

// SetName will store the user's name string into the session.
func (s Session) SetName(w http.ResponseWriter, r *http.Request, id string) error {
	return s.Put(w, r, UserNameKey, id)
}

// GetName retrieves the current user's name string from the session.
func (s Session) GetName(r *http.Request) string {
	name, err := s.Get(r, UserNameKey)
	if err != nil {
		return ""
	}

	return name
}

// SetEnvironment will store the user's environment string into the session.
func (s Session) SetEnvironment(w http.ResponseWriter, r *http.Request, id string) error {
	return s.Put(w, r, UserEnvironmentKey, id)
}

// GetEnvironment retrieves the current user's environment string from the session.
func (s Session) GetEnvironment(r *http.Request) string {
	environment, err := s.Get(r, UserEnvironmentKey)
	if err != nil {
		return ""
	}

	return environment
}
