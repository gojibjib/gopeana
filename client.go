// Europeana API wrapper
// https://pro.europeana.eu/resources/apis/search
// Inspired by https://github.com/nishanths/go-xkcd/
package gopeana

import (
	_ "fmt"
	"io"
	"net/http"
)

type Client struct {
	HTTPClient *http.Client
	Config
	apiKey string
}

func NewClient(apiKey string) *Client {
	return &Client{
		http.DefaultClient,
		Config{
			true,
		},
		apiKey,
	}
}

// baseUrl joins a constant url string with the client's API key and returns it as a string
func (c *Client) baseUrl() string {
	var protocol string
	const url = "www.europeana.eu/api/v2/search.json"

	if c.UseHTTPS {
		protocol = "https://"
	} else {
		protocol = "http://"
	}

	return protocol + url + "?wskey=" + c.apiKey
}

// do performs a basic HTTP request and returns the body
// User needs to make sure client is closed again
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
