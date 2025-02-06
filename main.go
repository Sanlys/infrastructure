package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"infrastructure/ipfetcher"
	"infrastructure/proxmox"
	"infrastructure/proxmox/authentication"
	"io"
	"net/http"
	"sync"
	"time"
)

type ProxmoxNode struct {
	Type   string `json:"type"`
	Status string `json:"status"`
	Disk   int    `json:"disk"`
	Node   string `json:"node"`
}

func fetchIP(index int, fetcher ipfetcher.IPFetcher, ip *string, ipChan chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d: Starting\n", index)
	fmt.Printf("Worker %d: Fetching IP address...\n", index)
	IP, err := fetcher.GetIP()
	if err != nil {
		fmt.Println(err.Error())
		close(ipChan)
		return
	}
	*ip = IP
	close(ipChan)
	fmt.Printf("Worker %d: Read IP %s\n", index, IP)
	fmt.Printf("Worker %d: Waiting...\n", index)
	time.Sleep(3 * time.Second)
	fmt.Printf("Worker %d: Finished waiting\n", index)
	fmt.Printf("Worker %d: Finished\n", index)
}

func useIP(index int, ip *string, ipChan chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d: Starting\n", index)

	<-ipChan
	output := *ip
	fmt.Printf("Worker %d: Received IP %s\n", index, output)
	fmt.Printf("Worker %d: Finished\n", index)
}

func onlyIP(index int, ip *string, ipChan chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d: Starting\n", index)

	<-ipChan
	output := *ip
	fmt.Printf("Worker %d: Received IP %s\n", index, output)
	fmt.Printf("Worker %d: Finished\n", index)
}

func testProxmoxConn(index int, endpoint string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d: Starting\n", index)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		fmt.Printf("Worker %d: Failed to create request: %s\n", index, err)
		return
	}

	req.Header.Set("Authorization", "")

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Worker %d: died %s\n", index, err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Printf("Worker %d: Status code: %d\n", index, res.StatusCode)
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Worker %d: Error reading response: %s\n", index, err)
		return
	}

	type ProxmoxResponse struct {
		Data []ProxmoxNode `json:"data"`
	}
	var response ProxmoxResponse

	jsonerr := json.Unmarshal(body, &response)
	if jsonerr != nil {
		fmt.Printf("Worker %d: Error unpacking json: %s\n", index, jsonerr)
	}

	nodes := response.Data

	formatted, _ := json.MarshalIndent(nodes, "", " ")
	fmt.Printf("Worker %d: Response: %s\n", index, string(formatted))
}

func main() {
	var wg sync.WaitGroup

	fetcher := ipfetcher.APIIPFetcher{}

	var (
		ip     string
		ipChan = make(chan struct{})
	)

	wg.Add(4)
	go fetchIP(1, fetcher, &ip, ipChan, &wg)
	go useIP(2, &ip, ipChan, &wg)
	go onlyIP(3, &ip, ipChan, &wg)
	go testProxmoxConn(4, "http://pve1.s1.lan:8006/api2/json/nodes", &wg)

	wg.Wait()
	fmt.Println("Finished")

	auth := authentication.InteractiveAuthentication{}

	proxmoxClient := proxmox.NewClient(
		"",
		&auth,
	)
}
