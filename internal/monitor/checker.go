package monitor

import (
	"net/http"
	"time"
)

func CheckURL(url string) (int, int64, bool) {

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	start := time.Now()

	resp, err := client.Get(url)

	if err != nil {
		return 0, 0, false
	}
	defer resp.Body.Close()
	responseTime := time.Since(start).Milliseconds()
	isUp := resp.StatusCode >= 200 && resp.StatusCode < 400
	return resp.StatusCode, responseTime, isUp
}
