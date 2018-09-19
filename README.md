# [gopeana](https://github.com/gojibjib/gopeana)
[![godoc badge](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/gojibjib/gopeana)
[![Go Report Card](https://goreportcard.com/badge/github.com/gojibjib/gopeana)](https://goreportcard.com/report/github.com/gojibjib/gopeana)
[![Build Status](https://travis-ci.org/gojibjib/gopeana.svg?branch=master)](https://travis-ci.org/gojibjib/gopeana)

An Europeana Search API client written in Go

[Europeana](https://www.europeana.eu) is a European collection of over 50 million digitised items.
The [Search API](https://pro.europeana.eu/resources/apis/search) provides a programmatic way to access those resources.
Make sure to [get an API key](https://pro.europeana.eu/get-api) first.

Inspired by [go-xkcd](https://github.com/nishanths/go-xkcd).

## Repo layout
The complete list of JibJib repos is:

- [jibjib](https://github.com/gojibjib/jibjib): Our Android app. Records sounds and looks fantastic.
- [deploy](https://github.com/gojibjib/deploy): Instructions to deploy the JibJib stack.
- [jibjib-model](https://github.com/gojibjib/jibjib-model): Code for training the machine learning model for bird classification
- [jibjib-api](https://github.com/gojibjib/jibjib-api): Main API to receive database requests & audio files.
- [jibjib-data](https://github.com/gojibjib/jibjib-data): A MongoDB instance holding information about detectable birds.
- [jibjib-query](https://github.com/gojibjib/jibjib-query): A thin Python Flask API that handles communication with the [TensorFlow Serving](https://www.tensorflow.org/serving/) instance.
- [gopeana](https://github.com/gojibjib/gopeana): A API client for [Europeana](https://europeana.eu), written in Go.
- [voice-grabber](https://github.com/gojibjib/voice-grabber): A collection of scripts to construct the dataset required for model training

## Install
```bash
$ go get github.com/gojibjib/gopeana
```

## Example
```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/gojibjib/gopeana"
	"log"
)

func main() {
	apiKey := "XXXXX"
	client := gopeana.NewClient(apiKey, "")
	   	
	// Returns all results for 'Mona Lisa' with an open license.
	request, err := gopeana.NewBasicSearchRequest(client, "open", "standard", "12", "1")
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
}
```