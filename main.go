package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Teachh/server-notify/internal/curler"
	"github.com/Teachh/server-notify/internal/logger"
	"github.com/Teachh/server-notify/internal/mail"
	"github.com/Teachh/server-notify/internal/telegram"
	"github.com/joho/godotenv"
)

var AVAILABE_METHODS = [3]string{"mail", "telegram", "whatsapp"}
var METHOD = ""

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		logger.Error.Fatalln("Error loading .env file")
	}
	METHOD = strings.ToLower(os.Args[1])
}

func main() {

	if contains(METHOD) {
		timePing, err := strconv.Atoi(os.Getenv("TIME_PING"))
		if err != nil {
			logger.Error.Println("Invalid TIME_PING value, defaulting to 1 minute")
			timePing = 1
		}
		for {
			sites := curler.GetCodes(strings.Split(os.Getenv("SITES"), ","))
			for site, code := range sites {
				if code != 200 {
					text := fmt.Sprintf("🔴 %s has code: %d", site, code)
					switch METHOD {
					case "mail":
						mail.SendMail(site, text)
					case "telegram":
						telegram.SendMessage(text)
					}
				} else {
					logger.Info.Printf("%s: %d", site, code)
				}
			}

			time.Sleep(time.Duration(timePing) * time.Minute)
		}
	} else {
		logger.Error.Fatalln("Method of sending messages not availabe, please use: mail or telegram")
	}
}

func contains(method string) bool {
	for _, v := range AVAILABE_METHODS {
		if v == method {
			return true
		}
	}
	return false
}
