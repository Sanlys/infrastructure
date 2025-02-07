package proxmox

import "fmt"

type ApiNodes ApiBase

type Node struct {
	Type            string  `json:"type"`
	Level           string  `json:"level"`
	Status          string  `json:"status"`
	Disk            int     `json:"disk"`
	Node            string  `json:"node"`
	SSL_Fingerprint string  `json:"ssl_fingerprint"`
	MaxMemory       int     `json:"maxmem"`
	Memory          int     `json:"mem"`
	MaxCPUS         int     `json:"maxcpu"`
	CPU             float64 `json:"cpu"`
	ID              string  `json:"id"`
	Uptime          int     `json:"uptime"`
}

type NodeResponse struct {
	Data []Node `json:"data"`
}

func (n *ApiNodes) GetNodes() ([]Node, error) {
	var response NodeResponse
	err := n.http.GET("/api2/json/nodes/", &response)
	return response.Data, err
}

type NodeDisk struct {
	Devpath   string   `json:"devpath"`
	GPT       int      `json:"gpt"`
	Mounted   bool     `json:"mounted"`
	OSDID     any      `json:"osdid"`
	OSDIDList []string `json:"osdid-list"`
	Size      int      `json:"size"`
}

type NodeDiskResponse struct {
	Data []NodeDisk `json:"data"`
}

func (n *ApiNodes) GetNodeDisks(nodeid string) ([]NodeDisk, error) {
	var response NodeDiskResponse
	err := n.http.GET(fmt.Sprintf("/api2/json/nodes/%s/disks/list", nodeid), &response)
	return response.Data, err
}
