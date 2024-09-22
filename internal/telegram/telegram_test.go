package telegram_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Teachh/server-notify/internal/telegram"
)

func TestSendMessage_Success(t *testing.T) {
	// Set required environment variables
	os.Setenv("TELEGRAM_TOKEN", "testToken")
	os.Setenv("TELEGRAM_CHAT_ID", "12345")

	// Mock HTTP server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]string
		json.NewDecoder(r.Body).Decode(&body)

		// Check if the right chat_id and text are being sent
		if body["chat_id"] != "12345" {
			t.Errorf("Expected chat_id to be 12345, got %s", body["chat_id"])
		}
		if body["text"] != "Test message" {
			t.Errorf("Expected text to be 'Test message', got %s", body["text"])
		}

		// Respond with a success status
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	// Send a test message
	telegram.SendMessage("Test message")

	// Clean up environment variables
	os.Unsetenv("TELEGRAM_TOKEN")
	os.Unsetenv("TELEGRAM_CHAT_ID")
}

func TestSendMessage_NoChatID(t *testing.T) {
	// Set only TELEGRAM_TOKEN and unset TELEGRAM_CHAT_ID
	os.Setenv("TELEGRAM_TOKEN", "testToken")
	os.Unsetenv("TELEGRAM_CHAT_ID")

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]string
		json.NewDecoder(r.Body).Decode(&body)

		// Check if the right chat_id and text are being sent
		if body["chat_id"] == "" {
			t.Errorf("Expected chat_id to be 12345, got %s", body["chat_id"])
			w.WriteHeader(http.StatusBadRequest)

		}
		if body["text"] != "Test message" {
			t.Errorf("Expected text to be 'Test message', got %s", body["text"])
		}

		// Respond with a success status
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	// This should cause a fatal error due to missing TELEGRAM_CHAT_ID
	telegram.SendMessage("Test message")
}
