package main

import (
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Custom testServer type which embeds a httptest.Server instance.
type testServer struct {
	*httptest.Server
 }

func newTestApplication(t *testing.T) *application {
	return &application{
		logger: slog.New(slog.DiscardHandler),
	}
}

func newTestServer(t *testing.T, h http.Handler) *testServer{
	ts := httptest.NewTLSServer(h)
	return &testServer{ts}
}

func (ts *testServer) get(t *testing.T, urlPath string) (int, http.Header, string) {
	rs, err := ts.Client().Get(urlPath)
	if err != nil {
		t.Fatal(err)
	}

	// And we can check that the response body written by the ping handler
	// equals "OK".
	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	return rs.StatusCode, rs.Header, string(body)
}
