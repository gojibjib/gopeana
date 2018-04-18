# [gopeana](https://github.com/gojibjib/gopeana)
An Europeana Search API client written in Go

[Europeana](https://www.europeana.eu) is a European collection of over 50 million digitised items.
The [Search API](https://pro.europeana.eu/resources/apis/search) provides a programmatic way to access those resources.
Make sure to [get an API key](https://pro.europeana.eu/get-api) first.

## Install
```bash
$ go get github.com/gojibjib/gopeana
```

## Example
Return all results for 'Mona Lisa' with an open license

```go
package main

   import (
   	"encoding/json"
   	"fmt"
   	"github.com/obitech/gopeana"
   	"log"
   )

   func main() {
   	apiKey := "XXXXX"
   	client := gopeana.NewClient(apiKey, "")
   	request, err := gopeana.NewRequest(client, "open", "", "", "")
   	if err != nil {
   		log.Fatal(err)
   	}
	
   	resp, err := request.Get("mona+lisa")
   	if err != nil {
   		log.Fatal(err)
   	}
   	
   	// Web search: https://www.europeana.eu/portal/de/search?q=mona+lisa&f%5BREUSABILITY%5D%5B%5D=open
	// API search: https://www.europeana.eu/api/v2/search.json?wskey=XXXXX&reusability=open&query=mona+lisa
   	fmt.Println(resp)

   	data, err := json.Marshal(resp)
   	if err != nil {
   		log.Fatalf("Error marshaling JSON")
   	}
   	fmt.Println(string(data))
   }
```