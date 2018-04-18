/*
Package gopeana provides an Europeana Search API client.

Example:

	apiKey := "XXXX"
	client := gopeana.NewClient(apiKey, "")

	// Will fetch results with open license, standard profile, default rows (12), default start (1)
	// Same as gopeana.NewRequest(client, "open", "", "", "")
	request, err := gopeana.NewRequest(client, "open", "standard", "12", "1")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := request.Get("mona+lisa")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", resp)

Europeana Search API documentation: https://pro.europeana.eu/resources/apis/search
*/
package gopeana
