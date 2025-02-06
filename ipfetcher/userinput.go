package ipfetcher

import "fmt"

type UserInputIPFetcher struct{}

func (u UserInputIPFetcher) GetIP() (string, error) {
	fmt.Print("Input IP: ")
	var input string
	fmt.Scan(&input)
	return input, nil
}
