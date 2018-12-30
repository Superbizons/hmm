package basic

var lastid = 0

type Client struct {
	ID      uint
	IP      string
	MaxBots int
}

func NewClient(ip string, maxbots int) *Client {
	lastid += 1
	return &Client{uint(lastid), ip, maxbots}
}
