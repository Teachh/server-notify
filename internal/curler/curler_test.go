package curler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Teachh/server-notify/internal/curler"
)

func TestGetCodes(t *testing.T) {
	expectedStatusCode := http.StatusOK
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(expectedStatusCode)
	}))
	defer mockServer.Close()
	urls := []string{mockServer.URL}
	result, err := curler.GetCodes(urls)
	if result[mockServer.URL] != expectedStatusCode || err != nil {
		t.Errorf("Expected status code %d, but got %d", expectedStatusCode, result[mockServer.URL])
	}
}
