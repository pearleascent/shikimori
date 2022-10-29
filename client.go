package shikimori

import (
	"net/http"
)

type Client struct {
	client *http.Client
}

func Init(client *http.Client) *Client {
	return &Client{
		client: client,
	}
}
