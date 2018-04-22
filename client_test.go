package gopeana

import (
	"fmt"
	"testing"
)

func assert(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got '%s', want '%s'", got, want)
	}
}

func TestClient(t *testing.T) {
	t.Run("Creating new Client with both keys", func(t *testing.T) {
		apiKey := "abc"
		privateKey := "def"
		client := NewClient(apiKey, privateKey)

		if client == nil {
			t.Errorf("Client is nil")
		}
	})

	t.Run("Creating new Client without private key", func(t *testing.T) {
		c := NewClient("abc", "")

		if c == nil {
			t.Errorf("Client is nil")
		}
	})

	t.Run("Return proper baseURL", func(t *testing.T) {
		c := NewClient("abc", "def")
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
