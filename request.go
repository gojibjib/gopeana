package gopeana

import (
	"fmt"
	"strconv"
)

// BasicSearchRequest is a wrapper around an Europeana API Search request, defining fields such
// as reusability, profile and rows/start for basic Pagination.
// You can pass an empty string to rows, profile or start to use the API default values
// rows = "" will return 12 results, start = "" will start with item 1, profile = "" will use standard profile.
type BasicSearchRequest struct {
	*Client
	reusability string
	profile     string
	rows        string
	start       string
}

type CursorSearchRequest struct {
	*Client
	reusability string
	profile     string
	cursor      string
}

// NewBasicSearchRequest returns a pointer to a BasicSearchRequest struct. This function will also perform error checking
// and return an error if an invalid value has been provided.
func NewBasicSearchRequest(c *Client, reusability, profile, rows, start string) (*BasicSearchRequest, error) {
	var request *BasicSearchRequest

	if err := checkReusability(reusability); err != nil {
		return request, err
	}

	if err := checkProfile(profile); err != nil {
		return request, err
	}

	if err := checkBasicPagination(rows, "rows can't be < 0", 0); err != nil {
		return request, err
	}

	if err := checkBasicPagination(start, "start can't be < 1", 1); err != nil {
		return request, err
	}

	return &BasicSearchRequest{
		Client:      c,
		reusability: reusability,
		profile:     profile,
		rows:        rows,
		start:       start,
	}, nil
}

// checkReusability will perform input validation for the reusability field
func checkReusability(s string) error {
	validReusability := []string{"", "open", "restricted", "permission"}
	for _, v := range validReusability {
		if s == v {
			return nil
		}
	}
	return fmt.Errorf("%s not part of valid arguments: %v",
		s, validReusability)
}

// checkProfile will perform input validation for the profile field
func checkProfile(s string) error {
	validProfile := []string{"", "minimal", "standard", "rich"}
	for _, v := range validProfile {
		if s == v {
			return nil
		}
	}
	return fmt.Errorf("%s not part of valid arguments: %s",
		s, validProfile)
}

// checkBasicPagination will will take a string check and try to convert it to an integer.
// If conversion fails or the converted value is smaller than a passed integer val,
// will return a custom error string passed as the info parameter.
// This function can be used to validate inputs for the rows and start field
func checkBasicPagination(check, info string, val int) error {
	if check != "" {
		check, err := strconv.Atoi(check)
		if err != nil {
			return err
		}
		if check < val {
			return fmt.Errorf("%s", info)
		}
		return nil
	}
	return nil
}

// searchUrl will use the struct's fields to construct a search URL and return it as string
func (r *BasicSearchRequest) searchURL() string {
	url := r.Client.baseURL()

	if r.reusability != "" {
		url += "&reusability=" + r.reusability
	}

	if r.profile != "" {
		url += "&profile=" + r.profile
	}

	if r.rows != "" {
		url += "&rows=" + r.rows
	}

	if r.start != "" {
		url += "&start=" + r.start
	}

	return url
}

// Reusability will set the reusability field or return an error
func (r *BasicSearchRequest) Reusability(s string) error {
	if err := checkReusability(s); err != nil {
		return err
	}

	r.reusability = s
	return nil
}

// Profile will set the profile field or return an error
func (r *BasicSearchRequest) Profile(s string) error {
	if err := checkProfile(s); err != nil {
		return err
	}
	r.profile = s
	return nil
}

// Rows will set the rows field or return an error
func (r *BasicSearchRequest) Rows(s string) error {
	if err := checkBasicPagination(s, "rows can't be < 0", 0); err != nil {
		return err
	}
	r.profile = s
	return nil
}

// Start will set the start field or return an error
func (r *BasicSearchRequest) Start(s string) error {
	if err := checkBasicPagination(s, "start can't be < 1", 1); err != nil {
		return err
	}
	r.profile = s
	return nil

}
