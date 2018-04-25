package gopeana

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var apiKey = os.Getenv("APIKEY")
var getRequests = []string{"mona+lisa", "tierstimmenarchiv"}

func checkErr(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Error(err)
	}
}

func TestApiKey(t *testing.T) {
	if apiKey == "" {
		t.Errorf("APIKEY is empty")
	}
}

func TestGetBasicSearchRequest(t *testing.T) {
	c := NewClient(apiKey, "")
	r, err := NewBasicSearchRequest(c, "", "", "", "")

	t.Run("Creating BasicSearchRequest", func(t *testing.T) {
		checkErr(t, err)
	})

	for _, query := range getRequests {
		t.Run(fmt.Sprintf("Running GET with '%s'", query), func(t *testing.T) {
			for _, re := range []string{"open"} {
				for _, pr := range []string{"minimal"} {
					for _, ro := range []string{""} {
						for _, st := range []string{""} {
							err := r.Reusability(re)
							checkErr(t, err)

							err = r.Profile(pr)
							checkErr(t, err)

							err = r.Rows(ro)
							checkErr(t, err)

							err = r.Start(st)
							checkErr(t, err)

							resp, err := Get(r, query)
							checkErr(t, err)

							if !resp.Success {
								t.Errorf("resp.Success should be true. resp.Error = %s", resp.Error)
							}

							if resp.Error != "" {
								t.Errorf("resp.Error should be empty. resp.Error = %s", resp.Error)
							}

							// Sleeping between requests because we're nice
							time.Sleep(500 * time.Millisecond)
						}
					}
				}
			}
		})
	}
}

func TestGetCursorSearchRequest(t *testing.T) {
	c := NewClient(apiKey, "")
	r, err := NewCursorSearchRequest(c, "", "", "")

	t.Run("Creating CursorSearchRequest", func(t *testing.T) {
		checkErr(t, err)
	})

	for _, query := range getRequests {
		t.Run(fmt.Sprintf("Running GET with '%s'", query), func(t *testing.T) {
			for _, re := range []string{"open"} {
				for _, pr := range []string{"minimal"} {
					for _, cu := range []string{""} {
						err := r.Reusability(re)
						checkErr(t, err)

						err = r.Profile(pr)
						checkErr(t, err)

						resp, err := Get(r, query)
						checkErr(t, err)

						r.Cursor(cu)

						if !resp.Success {
							t.Errorf("resp.Success should be true. resp.Error = %s", resp.Error)
						}

						if resp.Error != "" {
							t.Errorf("resp.Error should be empty. resp.Error = %s", resp.Error)
						}

						// Sleeping between requests because we're nice
						time.Sleep(500 * time.Millisecond)
					}
				}
			}
		})
	}
}
