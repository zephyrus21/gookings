package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

//! creates a new CSRF token tp all POST requests
func NoSurf(next http.Handler) http.Handler {
	//# creates a new CSRF handler
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

//! loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
