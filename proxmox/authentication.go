package proxmox

type Authentication interface {
	GetAuthHeader() (string, error)
}
