package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/Teachh/server-notify/internal/logger"
)

func generateUrl() string {
	token := os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		logger.Error.Fatalln("Add TELEGRAM_TOKEN in .env")
	}
	return fmt.Sprintf("https://api.telegram.org/bot%s", token)
}

// https://medium.com/geekculture/how-to-use-go-to-send-telegram-messages-to-your-phone-a819bdf7f35c
func SendMessage(text string) {
	chat_id := os.Getenv("TELEGRAM_CHAT_ID")
	if chat_id == "" {
		logger.Error.Println("Add TELEGRAM_CHAT_ID in .env")
		return
	}
	url := fmt.Sprintf("%s/sendMessage", generateUrl())
	body, _ := json.Marshal(map[string]string{
		"chat_id":    chat_id,
		"text":       text,
		"parse_mode": "markdown",
	})
	response, err := http.Post(
		url,
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		logger.Error.Fatal("Error sending message", err)
	}
	defer response.Body.Close()
	logger.Info.Println("Telegram Message sent:", text)
}
