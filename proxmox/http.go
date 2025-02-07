package proxmox

import (
	"encoding/json"
	"net/http"
)

type Http struct {
	baseurl     string
	auth        Authentication
	http_client http.Client
}

func (h *Http) GET(url string, data any) error {
	req, err := http.NewRequest("GET", h.baseurl+url, nil)
	if err != nil {
		return err
	}
	auth, err := h.auth.GetAuthHeader()
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", auth)

	resp, err := h.http_client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return err
	}

	return nil
}
