package client

import "net/http"

type Client struct {
	http *http.Client
}

type Request struct {
	Method string
	URL string
	Headers map[string]string
	Params map[string]string
}

type Response struct {
	StatusCode int
	Body string
}

func NewClient() *Client {
	return &Client{
		http: &http.Client{},
	}
}

func (c *Client) SendRequest() {

}
