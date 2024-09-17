package curler

import (
	"net/http"

	"github.com/Teachh/server-notify/internal/logger"
)

func GetCodes(urls []string) map[string]int {
	codes := make(map[string]int)
	for _, url := range urls {
		request, err := http.Get(url)
		if err != nil {
			logger.Error.Println("Error getting the status from", url, ". Error:", err)
		}
		codes[url] = request.StatusCode
	}
	return codes
}
