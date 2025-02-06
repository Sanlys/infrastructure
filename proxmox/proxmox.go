package proxmox

type Client struct {
	baseurl string
	auth    Authentication
	val     int

	// Endpoints
	//Nodes *nodes.NodesService
}

func NewClient(baseurl string, auth authentication.Authentication) *Client {
	if auth == nil {
		auth = &authentication.InteractiveAuthentication{}
	}
	return &Client{
		baseurl: baseurl,
		auth:    auth,
		val:     a,
	}
}
