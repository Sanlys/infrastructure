package proxmox

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Http struct {
	baseurl     string
	auth        Authentication
	http_client http.Client
}

type Metadata struct {
	Status        string
	StatusCode    int
	ContentLength int64
	Headers       http.Header
}

func (h *Http) GET(url string, data any) (Metadata, error) {
	metadata := Metadata{}
	req, err := http.NewRequest("GET", h.baseurl+url, nil)
	if err != nil {
		return metadata, err
	}
	auth, err := h.auth.GetAuthHeader()
	if err != nil {
		return metadata, err
	}
	req.Header.Set("Authorization", auth)

	resp, err := h.http_client.Do(req)
	if err != nil {
		return metadata, err
	}
	defer resp.Body.Close()
	metadata.StatusCode = resp.StatusCode
	metadata.Status = resp.Status
	metadata.ContentLength = resp.ContentLength
	metadata.Headers = resp.Header

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return metadata, err
	}

	return metadata, nil
}

func (h *Http) POST(url string, data any, input any) (Metadata, error) {
	metadata := Metadata{}

	body, err := json.Marshal(input)
	if err != nil {
		return metadata, err
	}

	req, err := http.NewRequest("POST", h.baseurl+url, bytes.NewBuffer(body))
	if err != nil {
		return metadata, err
	}
	auth, err := h.auth.GetAuthHeader()
	if err != nil {
		return metadata, err
	}
	req.Header.Set("Authorization", auth)
	req.Header.Set("Content-Type", "application/json")

	resp, err := h.http_client.Do(req)
	if err != nil {
		return metadata, err
	}
	defer resp.Body.Close()
	metadata.StatusCode = resp.StatusCode
	metadata.Status = resp.Status
	metadata.ContentLength = resp.ContentLength
	metadata.Headers = resp.Header

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return metadata, err
	}

	return metadata, nil
}
