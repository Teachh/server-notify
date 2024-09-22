package mail_test

import (
	"errors"
	"net/smtp"
	"os"
	"testing"

	"github.com/Teachh/server-notify/internal/mail"
)

func TestSendMail_Success(t *testing.T) {
	// Mock the SendMail function to simulate a successful email send
	mail.SendMailFunc = func(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
		// Simulate success
		return nil
	}

	// Set environment variables for testing
	os.Setenv("MAIL_FROM", "testsender@example.com")
	os.Setenv("MAIL_PASSWORD", "testpassword")
	os.Setenv("MAIL_TO", "testrecipient@example.com")

	// Call the actual SendMail function
	err := mail.SendMail("https://example.com", "The site is down")
	if err != nil {
		t.Errorf("SendMail failed: %v", err)
	}

	// Reset the SendMailFunc to the original smtp.SendMail after test
	mail.SendMailFunc = smtp.SendMail
}

func TestSendMail_Failure(t *testing.T) {
	// Mock the SendMail function to simulate an error in sending email
	mail.SendMailFunc = func(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
		// Simulate a failure
		return errors.New("failed to send email")
	}

	// Set environment variables for testing
	os.Setenv("MAIL_FROM", "testsender@example.com")
	os.Setenv("MAIL_PASSWORD", "testpassword")
	os.Setenv("MAIL_TO", "testrecipient@example.com")

	// Call the actual SendMail function
	err := mail.SendMail("https://example.com", "The site is down")
	if err == nil {
		t.Error("Expected SendMail to fail, but it succeeded")
	}

	// Reset the SendMailFunc to the original smtp.SendMail after test
	mail.SendMailFunc = smtp.SendMail
}
