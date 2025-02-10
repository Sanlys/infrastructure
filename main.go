package main

import (
	"fmt"
	"infrastructure/proxmox"
)

type ProxmoxNode struct {
	Type   string `json:"type"`
	Status string `json:"status"`
	Disk   int    `json:"disk"`
	Node   string `json:"node"`
}

func main() {

	auth := proxmox.InteractiveAuthentication{}

	proxmoxClient := proxmox.NewClient(
		"http://pve1.s1.lan:8006",
		&auth,
		nil,
	)
	/*nodes, err := proxmoxClient.Nodes.GetNodes()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, node := range nodes {
		fmt.Printf("\nNode: %s\n", node.Node)
		fmt.Printf("ID: %s\n", node.ID)
		fmt.Printf("Status: %s\n", node.Status)
		fmt.Printf("CPU: %d cores\n", node.MaxCPUS)
		fmt.Printf("Ram: %.2f GiB\n", float64(node.MaxMemory)/(1024*1024*1024))
		fmt.Printf("Uptime: %d\n", node.Uptime/(60*60*24))
	}
	disks, err := proxmoxClient.Nodes.GetNodeDisks("pve1")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, disk := range disks {
		fmt.Printf("\nMountpath: %s\n", disk.Devpath)
		fmt.Printf("Mounted: %t\n", disk.Mounted)
		fmt.Printf("Size: %.2f\n", float64(disk.Size)/(1000*1000*1000))
		fmt.Printf("GPT: %t\n", disk.GPT)
		fmt.Printf("OSD ID: %d\n", disk.OSDID)
		fmt.Printf("OSD ID List: %+v\n", disk.OSDIDList)
		fmt.Printf("Healht: %s\n", disk.Health)
		fmt.Printf("Model: %s\n", disk.Model)
		fmt.Printf("Parent: %s\n", disk.Parent)
		fmt.Printf("Serial: %s\n", disk.Serial)
		fmt.Printf("Used: %s\n", disk.Used)
		fmt.Printf("Vendor: %s\n", disk.Vendor)
		fmt.Printf("WWN: %s\n", disk.WWN)
	}*/

	vms, err := proxmoxClient.Nodes.GetQemuList("pve1")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, vm := range vms {
		fmt.Printf("\nVM ID: %d\n", vm.VMID)
		fmt.Printf("Status: %s\n", vm.Status)
		fmt.Printf("Template: %t\n", vm.Template)
		fmt.Printf("Name: %s\n", vm.Name)
		fmt.Printf("Tags: %s\n", vm.Tags)
		fmt.Printf("Uptime: %d\n", vm.Uptime)
		fmt.Printf("Lock: %s\n", vm.Lock)
		fmt.Printf("CPUs: %d\n", vm.CPUs)
		fmt.Printf("Max memory: %d\n", vm.MaxMem)
		fmt.Printf("Max disk: %d\n", vm.MaxDisk)
		fmt.Printf("Disk read: %d\n", vm.DiskRead)
		fmt.Printf("Disk write: %d\n", vm.DiskWrite)
		fmt.Printf("Net in: %d\n", vm.NetIn)
		fmt.Printf("Net out: %d\n", vm.NetOut)
		fmt.Printf("Qemu pid: %d\n", vm.Pid)
		fmt.Printf("QMP Status: %s\n", vm.QmpStatus)
		fmt.Printf("Qemu version: %s\n", vm.RunningQemu)
	}
}
