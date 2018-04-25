package gopeana

import (
	"os"
	"testing"
)

var apiKey = os.Getenv("APIKEY")

func TestApiKey(t *testing.T) {
	if apiKey == "" {
		t.Errorf("APIKEY is empty")
	}
}
