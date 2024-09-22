package mail

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/Teachh/server-notify/internal/logger"
)

var SendMailFunc = smtp.SendMail

func SendMail(webpage string, text string) error {
	auth := smtp.PlainAuth("", os.Getenv("MAIL_FROM"), os.Getenv("MAIL_PASSWORD"), "smtp.gmail.com")
	to := []string{os.Getenv("MAIL_TO")}
	msg := []byte(fmt.Sprintf("To: %s\r\n"+
		"Subject: %s down!\r\n"+
		"\r\n"+
		"%s\r\n", webpage, os.Getenv("MAIL_TO"), text))
	err := SendMailFunc("smtp.gmail.com:587", auth, os.Getenv("MAIL_FROM"), to, msg)
	if err != nil {
		logger.Error.Println("Error sending Mail", err)
		return err
	}
	return nil
}
