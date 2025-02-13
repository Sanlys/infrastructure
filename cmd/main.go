package main

import (
	"fmt"
	"infrastructure/proxmox"
	"strings"

	"github.com/rivo/tview"
)

type ProxmoxNode struct {
	Type   string `json:"type"`
	Status string `json:"status"`
	Disk   int    `json:"disk"`
	Node   string `json:"node"`
}

func HumanReadableDuration(seconds int64) string {
	if seconds < 0 {
		return "Invalid duration"
	}

	// Define the time units
	const (
		Second = 1
		Minute = 60 * Second
		Hour   = 60 * Minute
		Day    = 24 * Hour
	)

	units := []struct {
		Name   string
		Amount int64
	}{
		{"day", Day},
		{"hour", Hour},
		{"minute", Minute},
		{"second", Second},
	}

	var parts []string

	for _, unit := range units {
		if seconds >= unit.Amount {
			value := seconds / unit.Amount
			seconds %= unit.Amount
			unitName := unit.Name
			if value > 1 {
				unitName += "s"
			}
			parts = append(parts, fmt.Sprintf("%d %s", value, unitName))
		}
	}

	if len(parts) == 0 {
		return "0 seconds"
	}

	return strings.Join(parts, ", ")
}

func main() {

	auth := proxmox.InteractiveAuthentication{}

	proxmoxClient := proxmox.NewClient(
		"http://pve1.s1.lan:8006",
		&auth,
		nil,
	)

	nodes, err := proxmoxClient.Nodes.GetNodes()
	if err != nil {
		panic(err)
	}

	app := tview.NewApplication()
	list := tview.NewList()
	vmList := tview.NewList()
	infoBox := tview.NewTextView()
	flex := tview.NewFlex().
		AddItem(list, 0, 1, true).
		AddItem(vmList, 0, 1, false).
		AddItem(infoBox, 0, 3, false)

	for _, node := range nodes {
		list.AddItem(node.Node, fmt.Sprintf("ID: %s", node.ID), ' ', func() {
			infoBox.SetText(fmt.Sprintf("Node: %s\nID: %s\nUptime: %s", node.Node, node.ID, HumanReadableDuration(int64(node.Uptime))))

			vms, err := proxmoxClient.Nodes.GetQemuList(node.Node)
			if err != nil {
				panic(err)
			}
			vmList.Clear()
			for _, vm := range vms {
				vmList.AddItem(vm.Name, fmt.Sprintf("ID: %d", vm.VMID), ' ', func() {
					infoBox.SetText(fmt.Sprintf("VM: %s\nID: %d\nStatus: %s\nUptime: %s", vm.Name, vm.VMID, vm.Status, HumanReadableDuration(int64(node.Uptime))))
				})
			}
			app.SetFocus(vmList)
			vmList.AddItem("Return", "Return to nodes", 'r', func() {
				vmList.Clear()
				app.SetFocus(list)
			})
		})
	}

	list.AddItem("Quit", "Press to exit", 'q', func() {
		app.Stop()
	})
	app.SetRoot(flex, true)
	if err := app.Run(); err != nil {
		fmt.Printf("Error running application: %s\n", err)
	}
}
