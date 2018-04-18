package gopeana

import (
	"io"
	"net/http"
)

type Client struct {
	HTTPClient *http.Client
	Config
	ApiKey     string
	PrivateKey string
}

func NewClient(apiKey, privateKey string) *Client {
	return &Client{
		http.DefaultClient,
		Config{
			true,
		},
		apiKey,
		privateKey,
	}
}

// baseUrl joins a constant url string with the client's API key and returns it as a string.
func (c *Client) baseUrl() string {
	var protocol string
	const url = "www.europeana.eu/api/v2/search.json"

	if c.UseHTTPS {
		protocol = "https://"
	} else {
		protocol = "http://"
	}

	return protocol + url + "?wskey=" + c.ApiKey
}

// do performs a basic HTTP request and returns the body.
// User needs to make sure client is closed again.
func (c *Client) do(req *http.Request) (io.ReadCloser, error) {
	res, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, newStatusError(res.StatusCode)
	}

	return res.Body, nil
}
