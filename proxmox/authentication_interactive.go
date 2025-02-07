package proxmox

import (
	"sync"

	"github.com/manifoldco/promptui"
)

type InteractiveAuthentication struct {
	cache string
	mu    sync.Mutex
}

func (i *InteractiveAuthentication) GetAuthHeader() (string, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	if i.cache != "" {
		return i.cache, nil
	}
	usernamePrompt := promptui.Prompt{
		Label: "Username",
	}
	username, err := usernamePrompt.Run()
	if err != nil {
		return "", err
	}

	passwordPrompt := promptui.Prompt{
		Label: "Password",
		Mask:  '*',
	}
	password, err := passwordPrompt.Run()
	if err != nil {
		return "", err
	}
	i.cache = username + password
	return i.cache, nil
}
