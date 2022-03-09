package ch

import (
	"net/http"
	"strings"
)

const (
	acceptCH = "Accept-CH"
)

// Add Accept-CH to the response header.
//
// Note: Client hints are accessible only on secure origins (via TLS).
//
// Example:
// 	ch.Announce(w, []string{headers.SecChUa, headers.SecChUaBitness})
func Announce(w http.ResponseWriter, headers []string) {
	value := strings.Join(headers, ", ")

	w.Header().Add(acceptCH, value)
	w.Header().Add("Vary", value)
}

// Add Accept-CH to the response header.
// No Vary header is added.
//
// Note: Client hints are accessible only on secure origins (via TLS).
//
// Example:
// 	ch.AnnounceNoVary(w, []string{headers.SecChUa, headers.SecChUaBitness})
func AnnounceNoVary(w http.ResponseWriter, headers []string) {
	value := strings.Join(headers, ", ")

	w.Header().Add(acceptCH, value)
}
