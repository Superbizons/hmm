package basic

var lastid = 0

type Client struct {
	ID uint
	IP string
}

func NewClient(ip string) *Client {
	lastid += 1
	return &Client{uint(lastid), ip}
}
