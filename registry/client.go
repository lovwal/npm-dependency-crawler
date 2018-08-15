package registry

type Client struct {
	baseURL string
}

func NewClient(URL string) *Client {
	return &Client{
		baseURL: URL,
	}
}
