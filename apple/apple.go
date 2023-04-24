package apple

import (
	"net/http"
)

type Client struct {
	client *http.Client
}

func New() *Client {
	var c = &Client{}
	c.client = http.DefaultClient
	return c
}
