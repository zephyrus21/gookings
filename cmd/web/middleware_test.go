package main

import (
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var myH myHandler

	h := NoSurf(&myH)

	switch h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Errorf("NoSurf() returned a %T, want a http.Handler", h)
	}
}

func TestSessionLoad(t *testing.T) {
	var myH myHandler

	h := SessionLoad(&myH)

	switch v := h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Errorf("SessionLoad() returned a %T, want a http.Handler", v)
	}
}
