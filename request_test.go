package gopeana

import (
	"fmt"
	"testing"
)

var validReusability = []string{"", "open", "restricted", "permission"}
var validProfile = []string{"", "minimal", "standard", "rich"}
var validRows = []string{"", "0", "1", "12", "24"}
var validStart = []string{"", "1", "5", "18"}
var valiBasicRequests = []struct {
	input []string
	want  error
}{
	{[]string{"", "", "", ""}, nil},
	{[]string{"open", "", "", ""}, nil},
	{[]string{"restricted", "", "", ""}, nil},
	{[]string{"permission", "", "", ""}, nil},
	{[]string{"", "minimal", "", ""}, nil},
	{[]string{"", "standard", "", ""}, nil},
	{[]string{"", "rich", "", ""}, nil},
	{[]string{"", "", "0", "1"}, nil},
	{[]string{"", "", "12", "2"}, nil},
	{[]string{"open", "minimal", "12", "2"}, nil},
}
var validCursorRequests = []struct {
	input []string
	want  []error
}{
	{[]string{"", "", ""}, nil},
	{[]string{"", "", "*"}, nil},
	{[]string{"", "", "1235"}, nil},
	{[]string{"", "", "asp[el1230d"}, nil},
	{[]string{"", "", ""}, nil},
	{[]string{"open", "", ""}, nil},
	{[]string{"restricted", "", ""}, nil},
	{[]string{"permission", "", ""}, nil},
	{[]string{"", "minimal", ""}, nil},
	{[]string{"", "standard", ""}, nil},
	{[]string{"", "rich", ""}, nil},
	{[]string{"open", "minimal", "*"}, nil},
}

func assertURL(t *testing.T, c *Client, r *BasicSearchRequest, v, param string) {
	t.Helper()

	got := r.searchURL()
	want := ""
	if v == "" {
		want = c.baseURL()
	} else {
		want = fmt.Sprintf("%s&%s=%s", c.baseURL(), param, v)
	}
	if got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}
}

func TestValidNewRequest(t *testing.T) {
	c := NewClient("abc", "")

	t.Run("New BasicSearchRequest", func(t *testing.T) {
		for _, tt := range valiBasicRequests {
			if _, err := NewBasicSearchRequest(c, tt.input[0], tt.input[1], tt.input[2], tt.input[3]); err != nil {
				t.Errorf("error while creating new Request: %s", err)
			}
		}
	})

	t.Run("New CursorSearchRequest", func(t *testing.T) {
		for _, tt := range validCursorRequests {
			if _, err := NewCursorSearchRequest(c, tt.input[0], tt.input[1], tt.input[2]); err != nil {
				t.Errorf("error while creating new Request: %s", err)
			}
		}
	})

}

func TestValidSearchURL(t *testing.T) {
	c := NewClient("abc", "")

	t.Run("Basic URL", func(t *testing.T) {
		r, err := NewBasicSearchRequest(c, "", "", "", "")
		if err != nil {
			t.Errorf("%s", err)
		}

		got := r.searchURL()
		if got != c.baseURL() {
			t.Errorf("got: %s, want: %s", got, c.baseURL())
		}
	})

	t.Run("With Reusability", func(t *testing.T) {
		for _, v := range validReusability {
			r, err := NewBasicSearchRequest(c, v, "", "", "")
			if err != nil {
				t.Errorf("%s", err)
			}
			assertURL(t, c, r, v, "reusability")
		}
	})

	t.Run("With Profile", func(t *testing.T) {
		for _, v := range validProfile {
			r, err := NewBasicSearchRequest(c, "", v, "", "")
			if err != nil {
				t.Errorf("%s", err)
			}
			assertURL(t, c, r, v, "profile")
		}
	})

	t.Run("With Rows", func(t *testing.T) {
		for _, v := range validRows {
			r, err := NewBasicSearchRequest(c, "", "", v, "")
			if err != nil {
				t.Errorf("%s", err)
			}
			assertURL(t, c, r, v, "rows")
		}
	})

	t.Run("With Start", func(t *testing.T) {
		for _, v := range validStart {
			r, err := NewBasicSearchRequest(c, "", "", "", v)
			if err != nil {
				t.Errorf("%s", err)
			}
			assertURL(t, c, r, v, "start")
		}
	})

	t.Run("Full URL", func(t *testing.T) {
		for _, re := range validReusability[1:] {
			for _, p := range validProfile[1:] {
				for _, ro := range validRows[1:] {
					for _, s := range validStart[1:] {
						req, err := NewBasicSearchRequest(c, re, p, ro, s)
						if err != nil {
							t.Errorf("%s", err)
						}
						got := req.searchURL()
						want := fmt.Sprintf("%s&reusability=%s&profile=%s&rows=%s&start=%s", c.baseURL(), re, p, ro, s)
						if got != want {
							t.Errorf("got: %s, want: %s", got, want)
						}
					}
				}
			}
		}
	})
}

func TestInvalidNewBasicSearchRequest(t *testing.T) {
	c := NewClient("abc", "")

	t.Run("Invalid Reusability", func(t *testing.T) {
		for _, v := range []string{"abcd", "42", "closed", "How are you"} {
			if _, err := NewBasicSearchRequest(c, v, "", "", ""); err == nil {
				t.Errorf("error should have been thrown with reusability = %s", v)
			}
		}
	})

	t.Run("Invalid Profile", func(t *testing.T) {
		for _, v := range []string{"abcd", "42", "open", "How are you", "open"} {
			if _, err := NewBasicSearchRequest(c, "", v, "", ""); err == nil {
				t.Errorf("error should have been thrown with profile = %s", v)
			}
		}
	})

	t.Run("Invalid Rows", func(t *testing.T) {
		for _, v := range []string{"-1", "-15", "test", "xkcd", "43.2"} {
			if _, err := NewBasicSearchRequest(c, "", "", v, ""); err == nil {
				t.Errorf("error should have been thrown with rows = %s", v)
			}
		}
	})

	t.Run("Invalid Start", func(t *testing.T) {
		for _, v := range []string{"0", "-15", "test", "xkcd", "43.2"} {
			if _, err := NewBasicSearchRequest(c, "", "", "", v); err == nil {
				t.Errorf("error should have been thrown with rows = %s", v)
			}
		}
	})
}

func TestInvalidNewCursorSearchRequest(t *testing.T) {
	c := NewClient("abc", "")

	t.Run("Invalid Reusability", func(t *testing.T) {
		for _, v := range []string{"abcd", "42", "closed", "How are you"} {
			if _, err := NewCursorSearchRequest(c, v, "", "*"); err == nil {
				t.Errorf("error should have been thrown with reusability = %s", v)
			}
		}
	})

	t.Run("Invalid Profile", func(t *testing.T) {
		for _, v := range []string{"abcd", "42", "open", "How are you", "open"} {
			if _, err := NewCursorSearchRequest(c, "", v, "*"); err == nil {
				t.Errorf("error should have been thrown with profile = %s", v)
			}
		}
	})
}

func TestFieldChange(t *testing.T) {
	c := NewClient("abc", "")
	r, _ := NewBasicSearchRequest(c, "", "", "", "")

	t.Run("Change reusability", func(t *testing.T) {
		for _, v := range validReusability {
			if err := r.Reusability(v); err != nil {
				t.Errorf("%s", err)
			}
		}

		for _, v := range []string{"abc", "0123", "-15", "opent"} {
			if err := r.Reusability(v); err == nil {
				t.Errorf("err shouldn't be nil with r.Reusability(%s)", v)
			}
		}
	})

	t.Run("Change Profile", func(t *testing.T) {
		for _, v := range validProfile {
			if err := r.Profile(v); err != nil {
				t.Errorf("%s", err)
			}
		}

		for _, v := range []string{"abc", "0123", "-15", "standart"} {
			if err := r.Profile(v); err == nil {
				t.Errorf("err shouldn't be nil with r.Profile(%s)", v)
			}
		}
	})

	t.Run("Change Rows", func(t *testing.T) {
		for _, v := range validRows {
			if err := r.Rows(v); err != nil {
				t.Errorf("%s", err)
			}
		}

		for _, v := range []string{"-20", "3.14", "test", "-1"} {
			if err := r.Rows(v); err == nil {
				t.Errorf("err shouldn't be nil with r.Rows(%s)", v)
			}
		}
	})

	t.Run("Change Start", func(t *testing.T) {
		for _, v := range validStart {
			if err := r.Start(v); err != nil {
				t.Errorf("%s", err)
			}
		}

		for _, v := range []string{"-20", "3.14", "test", "0"} {
			if err := r.Start(v); err == nil {
				t.Errorf("err shouldn't be nil with r.Start(%s)", v)
			}
		}
	})
}
