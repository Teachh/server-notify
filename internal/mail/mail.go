package mail

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/Teachh/server-notify/internal/logger"
)

func SendMail(webpage string, text string) {
	auth := smtp.PlainAuth("", os.Getenv("MAIL_FROM"), os.Getenv("MAIL_PASSWORD"), "smtp.gmail.com")
	to := []string{os.Getenv("MAIL_TO")}
	msg := []byte(fmt.Sprintf("To: %s\r\n"+
		"Subject: %s down!\r\n"+
		"\r\n"+
		"%s\r\n", webpage, os.Getenv("MAIL_TO"), text))
	err := smtp.SendMail("smtp.gmail.com:587", auth, os.Getenv("MAIL_FROM"), to, msg)
	if err != nil {
		logger.Error.Fatal("Error sending Mail", err)
	}
}
