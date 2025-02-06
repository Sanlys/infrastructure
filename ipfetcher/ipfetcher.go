package ipfetcher

type IPFetcher interface {
	GetIP() (string, error)
}
