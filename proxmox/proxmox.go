package proxmox

import (
	"crypto/tls"
	"net/http"
)

type Client struct {
	baseurl     string
	auth        Authentication
	http_client http.Client

	// Endpoints
	Nodes ApiNodes
}

func NewClient(baseurl string, auth Authentication, http_client *http.Client) *Client {
	if auth == nil {
		auth = &InteractiveAuthentication{}
	}
	if http_client == nil {
		http_client = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		}

	}
	return &Client{
		baseurl:     baseurl,
		auth:        auth,
		http_client: *http_client,
		Nodes: ApiNodes{
			baseurl:     baseurl,
			auth:        auth,
			http_client: *http_client,
		},
	}
}
