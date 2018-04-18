package gopeana

import (
	"fmt"
	"strconv"
)

// SearchRequest is a wrapper around an Europeana API Search request, defining fields such
// as Reusability, Profile and Rows/Start for basic Pagination.
// You can pass an empty string to rows, profile or start to use the API default values
// rows = "" will return 12 results, start = "" will start with item 1, profile = "" will use standard profile.
type SearchRequest struct {
	Client      *Client
	Reusability string
	Profile     string
	Rows        string
	Start       string
}

// NewRequest returns a pointer to a SearchRequest struct. This function will also perform error checking
// and return an error if an invalid value has been provided.
func NewRequest(c *Client, reusability, profile, rows, start string) (*SearchRequest, error) {
	var request *SearchRequest

	validReusability := []string{"", "open", "restricted", "permission"}
	_, err := func() (bool, error) {
		for _, v := range validReusability {
			if reusability == v {
				return true, nil
			}
		}
		return false, fmt.Errorf("%s not part of valid arguments: %v",
			reusability, validReusability)
	}()
	if err != nil {
		return request, err
	}

	validProfile := []string{"", "minimal", "standard", "rich"}
	_, err = func() (bool, error) {
		for _, v := range validProfile {
			if profile == v {
				return true, nil
			}
		}
		return false, fmt.Errorf("%s not part of valid arguments: %s",
			profile, validProfile)
	}()
	if err != nil {
		return request, err
	}

	if rows != "" {
		check, err := strconv.Atoi(rows)
		if err != nil {
			return request, err
		}
		if check < 0 {
			return request, fmt.Errorf("rows can't be < 0")
		}
	}

	if start != "" {
		check, err := strconv.Atoi(start)
		if err != nil {
			return request, err
		}
		if check < 1 {
			return request, fmt.Errorf("start can't be < 1")
		}
	}

	return &SearchRequest{
		Client:      c,
		Reusability: reusability,
		Profile:     profile,
		Rows:        rows,
		Start:       start,
	}, nil
}

// searchUrl will use the struct's fields to construct a search URL and return it as string
func (r *SearchRequest) searchURL() string {
	url := r.Client.baseURL()

	if r.Reusability != "" {
		url += "&reusability=" + r.Reusability
	}

	if r.Profile != "" {
		url += "&profile=" + r.Profile
	}

	if r.Rows != "" {
		url += "&rows=" + r.Rows
	}

	if r.Start != "" {
		url += "&start=" + r.Start
	}

	return url
}
