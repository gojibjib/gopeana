package gopeana

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Item describes 'rich' metadata set described at
// https://pro.europeana.eu/resources/apis/search#profile-rich
type Item struct {
	ID                   string                `json:"id"`
	Title                []string              `json:"title"`
	TitleLang            map[string][]string   `json:"dcTitleLangAware"`
	Description          []string              `json:"dcDescription"`
	DescriptionLang      map[string][]string   `json:"dcDescriptionLangAware"`
	Completeness         int                   `json:"europeanaCompleteness"`
	DataProvider         []string              `json:"dataProvider"`
	Rights               []string              `json:"rights"`
	Source               []string              `json:"edmIsShownAt"`
	Latitude             []string              `json:"edmPlaceLatitude"`
	Longitude            []string              `json:"edmPlaceLongitude"`
	Preview              []string              `json:"edmPreview"`
	GUID                 string                `json:"guid"`
	Link                 string                `json:"link"`
	Type                 string                `json:"type"`
	Provider             []string              `json:"provider"`
	Creator              []string              `json:"dcCreator"`
	CreatorLang          map[string][]string   `json:"dcCreatorLangAware"`
	Score                int                   `json:"score"`
	Year                 []string              `json:"year"`
	ConceptTerm          []string              `json:"edmConceptTerm"`
	ConceptPrefLabel     []map[string][]string `json:"edmConceptPrefLabel"`
	ConceptPrefLabelLang map[string][]string   `json:"edmConceptPrefLabelLangAware"`
	ConceptBroaderTerm   []map[string][]string `json:"edmConceptBroaderTerm"`
	ConceptBroaderLabel  []map[string][]string `json:"edmConceptBroaderLabel"`
	TimespanLabel        []map[string]string   `json:"edmTimespanLabel"`
	TimespanLabelLang    map[string][]string   `json:"edmTimespanLabelLangAware"`
	Ugc                  []bool                `json:"ugc"`
	Country              []string              `json:"country"`
	DatasetName          []string              `json:"edmDatasetName"`
	Language             []string              `json:"dcLanguage"`
	TermIsPartOf         []string              `json:"dctermIsPartOf"`
	Timestamp            int                   `json:"timestamp"`
	TimestampCreated     string                `json:"timestampCreated"`
	TimestampUpdate      string                `json:"timestampUpdate"`
	IsShownBy            []string              `json:"edmIsShownBy"`
}

// Response describes a standard Europeana API response.
// apiKey has been omitted, since it can be gathered from the client.
type Response struct {
	Success       bool   `json:"success"`
	RequestNumber int    `json:"requestNumber"`
	ItemsCount    int    `json:"itemsCount"`
	TotalResults  int    `json:"totalResults"`
	NextCursor    string `json:"nextCursor"`
	Error         string `json:"error"`
	Items         []Item `json:"items"`
}

// doSearchRequest will perform a Europeana Search API request and return it as a Response struct.
// This will also close the body.
func doSearchRequest(r Request, query string) (Response, error) {
	var resp Response
	requestString := r.searchURL() + "&query=" + query

	log.Printf("Sending: GET %s", requestString)

	req, err := http.NewRequest("GET", requestString, nil)
	if err != nil {
		return resp, err
	}

	body, err := r.Client().do(req)
	if err != nil {
		return resp, err
	}
	defer body.Close()

	// Using json.Decoder here since we're reading from a stream
	if err := json.NewDecoder(body).Decode(&resp); err != nil {
		if !resp.Success {
			return resp, fmt.Errorf("%s", resp.Error)
		}
		return resp, err
	}

	return resp, nil
}

// Get returns an Europeana Search API response for the passed query and request.
func Get(r Request, query string) (Response, error) {
	return doSearchRequest(r, query)
}
