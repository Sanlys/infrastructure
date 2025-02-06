package proxmox

type Authentication interface {
	GetAuthHeader() (string, error)
}

const a = 1
