package ipfetcher

import (
	"bufio"
	"errors"
	"net/http"
)

type APIIPFetcher struct{}

func (a APIIPFetcher) GetIP() (string, error) {
	url := "https://api.ipify.org"

	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", errors.New("failed to fetch IP: received non-200 status code")
	}

	// Read response body
	scanner := bufio.NewScanner(res.Body)
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return "", err
		}
		return "", errors.New("failed to read IP from response")
	}

	ip := scanner.Text()
	return ip, nil
}
