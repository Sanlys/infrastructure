package proxmox

import (
	"encoding/json"
	"fmt"
	"strconv"
)

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
	GPT       bool     `json:"gpt"`
	Mounted   bool     `json:"mounted"`
	OSDID     int      `json:"osdid"`
	OSDIDList []string `json:"osdid-list"`
	Size      int      `json:"size"`

	// Optional
	Health string `json:"health"`
	Model  string `json:"model"`
	Parent string `json:"parent"`
	Serial string `json:"serial"`
	Used   string `json:"used"`
	Vendor string `json:"vendor"`
	WWN    string `json:"wwn"`
}

func (n *NodeDisk) UnmarshalJSON(data []byte) error {
	type Alias NodeDisk
	aux := &struct {
		OSDID any `json:"osdid"`
		GPT   any `json:"gpt"`
		*Alias
	}{
		Alias: (*Alias)(n),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	switch v := aux.OSDID.(type) {
	case float64:
		n.OSDID = int(v)
	case string:
		i, err := strconv.Atoi(v)
		if err != nil {
			n.OSDID = -1
		} else {
			n.OSDID = i
		}
	default:
		n.OSDID = -1
	}

	switch v := aux.GPT.(type) {
	case float64:
		n.GPT = v == 1
	case int:
		n.GPT = v == 1
	case string:
		i, err := strconv.Atoi(v)
		if err != nil {
			n.GPT = false
		} else {
			n.GPT = i == 1
		}
	default:
		n.GPT = false
	}

	return nil
}

type NodeDiskResponse struct {
	Data []NodeDisk `json:"data"`
}

func (n *ApiNodes) GetNodeDisks(nodeid string) ([]NodeDisk, error) {
	var response NodeDiskResponse
	err := n.http.GET(fmt.Sprintf("/api2/json/nodes/%s/disks/list", nodeid), &response)
	return response.Data, err
}

type Qemu struct {
	Status string `json:"status"`
	VMID   int    `json:"vmid"`

	// Optional
	CPUs           int    `json:"cpus"`
	DiskRead       int    `json:"diskread"`
	DiskWrite      int    `json:"diskwrite"`
	Lock           string `json:"lock"`
	MaxDisk        int    `json:"maxdisk"`
	MaxMem         int    `json:"maxmem"`
	Name           string `json:"name"`
	NetIn          int    `json:"netin"`
	NetOut         int    `json:"netout"`
	Pid            int    `json:"pid"`
	QmpStatus      string `json:"qmpstatus"`
	RunningMachine string `json:"running-machine"`
	RunningQemu    string `json:"running-qemu"`
	Tags           string `json:"tags"`
	Template       bool   `json:"template"`
	Uptime         int    `json:"uptime"`
}

func (n *Qemu) UnmarshalJSON(data []byte) error {
	type Alias Qemu
	aux := &struct {
		Template any `json:"template"`
		*Alias
	}{
		Alias: (*Alias)(n),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	switch v := aux.Template.(type) {
	case float64:
		n.Template = v == 1
	case int:
		n.Template = v == 1
	case string:
		i, err := strconv.Atoi(v)
		if err != nil {
			n.Template = false
		} else {
			n.Template = i == 1
		}
	default:
		n.Template = false
	}

	return nil
}

type QemuResponse struct {
	Data []Qemu `json:"data"`
}

func (n *ApiNodes) GetQemuList(nodeid string) ([]Qemu, error) {
	var response QemuResponse
	err := n.http.GET(fmt.Sprintf("/api2/json/nodes/%s/qemu", nodeid), &response)
	return response.Data, err
}
