package gopeana

import (
	"fmt"
	"testing"
)

func getClient(t *testing.T, apiKey, privateKey string) *Client {
	t.Helper()
	return NewClient(apiKey, privateKey)
}

func assert(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got '%s', want '%s'", got, want)
	}
}

func TestClient(t *testing.T) {
	t.Run("Creating new Client", func(t *testing.T) {
		apiKey := "abc"
		privateKey := "def"
		client := NewClient(apiKey, privateKey)

		if client == nil {
			t.Errorf("Client is nil")
		}
	})

	t.Run("Return proper baseURL", func(t *testing.T) {
		c := getClient(t, "abc", "def")
		got := c.baseURL()
		if c.Config.UseHTTPS {
			want := fmt.Sprintf("https://www.europeana.eu/api/v2/search.json?wskey=%s", c.APIKey)
			assert(t, got, want)
		} else {
			want := fmt.Sprintf("http://www.europeana.eu/api/v2/search.json?wskey=%s", c.APIKey)
			assert(t, got, want)
		}
	})
}
