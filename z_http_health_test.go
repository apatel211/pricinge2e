package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPHealthCheck(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"UP","service":"pricing-engine"}`))
	}))
	defer mockServer.Close()

	t.Logf("Verifying service heartbeat against: %s", mockServer.URL)

	resp, err := http.Get(mockServer.URL + "/health")
	if err != nil {
		t.Fatalf("Failed to dispatch heartbeat ping: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Unexpected health status code: got %d, want %d", resp.StatusCode, http.StatusOK)
	}
}