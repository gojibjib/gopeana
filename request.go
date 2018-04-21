package gopeana

import (
	"fmt"
	"strconv"
)

// SearchRequest is a wrapper around an Europeana API Search request, defining fields such
// as reusability, profile and rows/start for basic Pagination.
// You can pass an empty string to rows, profile or start to use the API default values
// rows = "" will return 12 results, start = "" will start with item 1, profile = "" will use standard profile.
type SearchRequest struct {
	*Client
	reusability string
	profile     string
	rows        string
	start       string
}

// NewRequest returns a pointer to a SearchRequest struct. This function will also perform error checking
// and return an error if an invalid value has been provided.
func NewRequest(c *Client, reusability, profile, rows, start string) (*SearchRequest, error) {
	var request *SearchRequest

	if err := checkReusability(reusability); err != nil {
		return request, err
	}

	if err := checkProfile(profile); err != nil {
		return request, err
	}

	if err := checkPagination(rows, "rows can't be < 0", 0); err != nil {
		return request, err
	}

	if err := checkPagination(start, "start can't be < 1", 1); err != nil {
		return request, err
	}

	return &SearchRequest{
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

// checkPagination will will take a string check and try to convert it to an integer.
// If conversion fails or the converted value is smaller than a passed integer val,
// will return a custom error string passed as the info parameter.
// This function can be used to validate inputs for the rows and start field
func checkPagination(check, info string, val int) error {
	if check != "" {
		check, err := strconv.Atoi(check)
		if err != nil {
			return err
		}
		if check < val {
			return fmt.Errorf("%s", info)
		}
	}
	return nil
}

// searchUrl will use the struct's fields to construct a search URL and return it as string
func (r *SearchRequest) searchURL() string {
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
func (r *SearchRequest) Reusability(s string) error {
	if err := checkReusability(s); err != nil {
		r.reusability = s
		return nil
	} else {
		return err
	}
}

// Reusability will set the profile or return an error
func (r *SearchRequest) Profile(s string) error {
	if err := checkProfile(s); err != nil {
		r.profile = s
		return nil
	} else {
		return err
	}
}

// Reusability will set the rows field or return an error
func (r *SearchRequest) Rows(s string) error {
	if err := checkPagination(s, "rows can't be < 0", 0); err != nil {
		r.profile = s
		return nil
	} else {
		return err
	}
}

// Reusability will set the start field or return an error
func (r *SearchRequest) Start(s string) error {
	if err := checkPagination(s, "start can't be < 0", 1); err != nil {
		r.profile = s
		return nil
	} else {
		return err
	}
}
