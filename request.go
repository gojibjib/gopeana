package gopeana

import (
	"errors"
	"fmt"
	"strconv"
)

type SearchRequest struct {
	Client      *Client
	Reusability string
	Profile     string
	Rows        string
	Start       string
}

func NewRequest(c *Client, reusability, profile, rows, start string) (*SearchRequest, error) {
	var request *SearchRequest

	validReusability := []string{"", "open", "restricted", "permission"}
	_, err := func() (bool, error) {
		for _, v := range validReusability {
			if reusability == v {
				return true, nil
			}
		}
		return false, errors.New(fmt.Sprintf("%s not part of valid arguments: %s",
			reusability, validReusability))
	}()
	if err != nil {
		return request, err
	}

	validProfile := []string{"", "minimal"}
	_, err = func() (bool, error) {
		for _, v := range validProfile {
			if profile == v {
				return true, nil
			}
		}
		return false, errors.New(fmt.Sprintf("%s not part of valid arguments: %s",
			profile, validProfile))
	}()
	if err != nil {
		return request, err
	}
	// Right now only minimal profile is supported
	if profile == "" {
		profile = "minimal"
	}

	if rows != "" {
		check, err := strconv.Atoi(rows)
		if err != nil {
			return request, err
		}
		if check < 0 {
			return request, errors.New("rows can't be < 0")
		}
	}

	if start != "" {
		check, err := strconv.Atoi(start)
		if err != nil {
			return request, err
		}
		if check < 1 {
			return request, errors.New("start can't be < 1")
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

func (r *SearchRequest) searchUrl() string {
	url := r.Client.baseUrl()

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
