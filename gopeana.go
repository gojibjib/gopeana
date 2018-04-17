package gopeana

import (
	"encoding/json"
	"net/http"
)

// Item defines the minimal profile
type Item struct {
	Id           string              `json:"id"`
	Title        []string            `json:"title"`
	Description  []string            `json:"dcDescription"`
	Completeness int                 `json:"europeanaCompleteness"`
	DataProvider []string            `json:"dataProvider"`
	Rights       []string            `json:"rights"`
	Source       []string            `json:"edmIsShownAt"`
	Latitude     []string            `json:"edmPlaceLatitude"`
	Longitude    []string            `json:"edmPlaceLongitude"`
	Preview      []string            `json:"edmPreview"`
	GUID         string              `json:"guid"`
	Link         string              `json:"link"`
	Type         string              `json:"type"`
	Provider     []string            `json:"provider"`
	Creator      []string            `json:"dcCreator"`
	CreatorLang  map[string][]string `json:"dcCreatorLangAware"`
	Score        int                 `json:"score"`
	Year         int                 `json:"year"`
}

type Response struct {
	apiKey        string `json:"-"`
	Success       bool   `json:"success"`
	RequestNumber int    `json:"requestNumber"`
	ItemsCount    int    `json:"itemsCount"`
	TotalResults  int    `json:"totalResults"`
	Items         []Item `json:"items"`
}

func (r *SearchRequest) doSearchRequest(query string) (Response, error) {
	var resp Response

	req, err := http.NewRequest("GET", r.searchUrl()+"&query="+query, nil)
	if err != nil {
		return resp, err
	}

	body, err := r.Client.do(req)
	if err != nil {
		return resp, err
	}
	defer body.Close()

	if err := json.NewDecoder(body).Decode(&resp); err != nil {
		return resp, err
	}

	return resp, nil
}

// Get returns an Europeana Search API response for the passed query
func (r *SearchRequest) Get(query string) (Response, error) {
	return r.doSearchRequest(query)
}
