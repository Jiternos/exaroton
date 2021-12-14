package exaroton

import (
	"net/http"
	"time"
)

// Exaroton API Wrapper Version
const VERSION = "1.0.2"

// Creates a new Exaroton session
func New(token string) (s *Session) {

	// Create a session interface
	s = &Session{
		Token:     "Bearer " + token,
		Client:    &http.Client{Timeout: (20 * time.Second)},
		UserAgent: "go-exaroton-api@" + VERSION,
	}

	return
}
