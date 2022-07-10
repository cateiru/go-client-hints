package ch_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cateiru/go-client-hints/ch"
	"github.com/cateiru/go-client-hints/ch/headers"
	"github.com/stretchr/testify/require"
)

type server struct {
	mux *http.ServeMux
}

func (s *server) announceHandler(w http.ResponseWriter, r *http.Request) {
	ch.Announce(w, []string{headers.SecChUa, headers.SecChUaMobile})
}

func (s *server) announceNoVaryHandler(w http.ResponseWriter, r *http.Request) {
	ch.AnnounceNoVary(w, []string{headers.SecChUa, headers.SecChUaMobile})
}

func new() *server {
	mux := http.NewServeMux()
	s := &server{mux: mux}

	s.mux.HandleFunc("/", s.announceHandler)
	s.mux.HandleFunc("/n", s.announceNoVaryHandler)

	return s
}

func TestAnnounce(t *testing.T) {
	s := httptest.NewServer(new().mux)

	resp, err := http.Get(s.URL + "/")
	require.NoError(t, err)

	h := resp.Header.Get("Accept-CH")
	require.Equal(t, h, "Sec-Ch-Ua, Sec-Ch-Ua-Mobile")

	v := resp.Header.Get("Vary")
	require.Equal(t, v, "Sec-Ch-Ua, Sec-Ch-Ua-Mobile")
}

func TestAnnounceNoVary(t *testing.T) {
	s := httptest.NewServer(new().mux)

	resp, err := http.Get(s.URL + "/n")
	require.NoError(t, err)

	h := resp.Header.Get("Accept-CH")
	require.Equal(t, h, "Sec-Ch-Ua, Sec-Ch-Ua-Mobile")

	v := resp.Header.Get("Vary")
	require.Empty(t, v)

}
