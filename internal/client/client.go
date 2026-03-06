package client

import (
	"io"
	"net/http"
	"strings"
)

type Client struct {
	http *http.Client
}

type Request struct {
	Method string
	URL string
	Headers map[string]string
}

type Response struct {
	StatusCode int
	Header http.Header
	Body string
}

func NewClient() *Client {
	return &Client{
		http: &http.Client{},
	}
}

func (c *Client) SendRequest(request Request) (*Response, error) {
	url := request.URL
	method := strings.TrimSpace(request.Method)
	// headers := request.Headers

	var resp *http.Response
	var err error

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}

	resp, err = c.http.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	strBody := string(body)
	
	response := &Response {
		StatusCode: resp.StatusCode,
		Body: strBody,
		Header: resp.Header,
	}

	return response, err

}
