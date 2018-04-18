package main

import (
	"encoding/json"
	"fmt"
	"github.com/obitech/gopeana"
	"log"
)

func main() {
	apiKey := "XXXX"
	client := gopeana.NewClient(apiKey, "")

	request, err := gopeana.NewRequest(client, "open", "", "", "")
	if err != nil {
		log.Fatal(err)
	}

	// Fetch all results for 'Mona Lisa' with an open license
	resp, err := request.Get("mona+lisa")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)

	// Print data as JSON
	data, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error marshaling JSON")
	}
	fmt.Println(string(data))
}
