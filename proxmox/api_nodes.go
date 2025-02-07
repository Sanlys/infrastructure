package proxmox

import (
	"encoding/json"
	"net/http"
)

type Node struct {
	Type   string `json:"type"`
	Status string `json:"status"`
	Disk   int    `json:"disk"`
	Node   string `json:"node"`
	//SSL_Fingerprint string `json:"ssl_fingerprint"`
	Memory int    `json:"mem"`
	CPUs   int    `json:"maxcpu"`
	ID     string `json:"id"`
	Uptime int    `json:"uptime"`
}

type NodeResponse struct {
	Data []Node `json:"data"`
}

type ApiNodes struct { // Refactor
	http_client http.Client
	baseurl     string
	auth        Authentication
}

func (n *ApiNodes) GetNodes() ([]Node, error) {
	req, err := http.NewRequest("GET", n.baseurl+"/api2/json/nodes", nil)
	if err != nil {
		return nil, err
	}
	auth, err := n.auth.GetAuthHeader()
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", auth)

	resp, err := n.http_client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response NodeResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}
