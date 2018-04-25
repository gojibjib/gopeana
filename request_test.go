package gopeana

import (
	"fmt"
	"testing"
)

var validReusability = []string{"", "open", "restricted", "permission"}
var validProfile = []string{"", "minimal", "standard", "rich"}
var validRows = []string{"", "0", "1", "12", "24"}
var validStart = []string{"", "1", "5", "18"}
var validCursor = []string{"", "*", "123asd", "12as123sd"}
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

func assertURL(t *testing.T, c *Client, r Request, key, wantVal string) {
	t.Helper()

	got := r.searchURL()
	want := ""
	if wantVal == "" {
		want = c.baseURL()
	} else {
		want = fmt.Sprintf("%s&%s=%s", c.baseURL(), key, wantVal)
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

func TestClientFunction(t *testing.T) {
	c := NewClient("abc", "")

	t.Run("BasicSearchRequest.Client()", func(t *testing.T) {
		r, err := NewBasicSearchRequest(c, "", "", "", "")
		if err != nil {
			t.Error(err)
		}

		if r.Client() != c {
			t.Errorf("r.Client(): got %v, want %v", r.Client(), c)
		}
	})

	t.Run("CursorSearchRequest.Client()", func(t *testing.T) {
		r, err := NewCursorSearchRequest(c, "", "", "")
		if err != nil {
			t.Error(err)
		}

		if r.Client() != c {
			t.Errorf("r.Client(): got %v, want %v", r.Client(), c)
		}
	})
}

func TestValidBasicSearchURL(t *testing.T) {
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
			assertURL(t, c, r, "reusability", v)
		}
	})

	t.Run("With Profile", func(t *testing.T) {
		for _, v := range validProfile {
			r, err := NewBasicSearchRequest(c, "", v, "", "")
			if err != nil {
				t.Errorf("%s", err)
			}
			assertURL(t, c, r, "profile", v)
		}
	})

	t.Run("With Rows", func(t *testing.T) {
		for _, v := range validRows {
			r, err := NewBasicSearchRequest(c, "", "", v, "")
			if err != nil {
				t.Errorf("%s", err)
			}
			assertURL(t, c, r, "rows", v)
		}
	})

	t.Run("With Start", func(t *testing.T) {
		for _, v := range validStart {
			r, err := NewBasicSearchRequest(c, "", "", "", v)
			if err != nil {
				t.Errorf("%s", err)
			}
			assertURL(t, c, r, "start", v)
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

func TestValidCursorSearchURL(t *testing.T) {
	c := NewClient("abc", "")

	t.Run("Basic URL", func(t *testing.T) {
		req, err := NewCursorSearchRequest(c, "", "", "")
		if err != nil {
			t.Errorf("%s", err)
		}
		assertURL(t, c, req, "cursor", "*")
	})

	t.Run("With Reusability", func(t *testing.T) {
		for _, v := range validReusability {
			r, err := NewCursorSearchRequest(c, v, "", "")
			if err != nil {
				t.Errorf("%s", err)
			}
			got := r.searchURL()
			want := ""
			if v == "" {
				want = fmt.Sprintf("%s&cursor=%s", c.baseURL(), "*")
			} else {
				want = fmt.Sprintf("%s&reusability=%s&cursor=%s", c.baseURL(), v, "*")
			}
			if got != want {
				t.Errorf("got: %s, want: %s", got, want)
			}
		}
	})

	t.Run("With Profile", func(t *testing.T) {
		for _, v := range validProfile {
			r, err := NewCursorSearchRequest(c, "", v, "")
			if err != nil {
				t.Errorf("%s", err)
			}
			got := r.searchURL()
			want := ""
			if v == "" {
				want = fmt.Sprintf("%s&cursor=%s", c.baseURL(), "*")
			} else {
				want = fmt.Sprintf("%s&profile=%s&cursor=%s", c.baseURL(), v, "*")
			}
			if got != want {
				t.Errorf("got: %s, want: %s", got, want)
			}
		}
	})

	t.Run("Full URL", func(t *testing.T) {
		for _, re := range validReusability[1:] {
			for _, p := range validProfile[1:] {
				for _, cu := range validCursor[1:] {
					req, err := NewCursorSearchRequest(c, re, p, cu)
					if err != nil {
						t.Errorf("%s", err)
					}
					got := req.searchURL()
					want := fmt.Sprintf("%s&reusability=%s&profile=%s&cursor=%s", c.baseURL(), re, p, cu)
					if got != want {
						t.Errorf("got: %s, want: %s", got, want)
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

func TestFieldChangeBasicSearchRequest(t *testing.T) {
	c := NewClient("abc", "")
	r, _ := NewBasicSearchRequest(c, "", "", "", "")

	t.Run("Change reusability", func(t *testing.T) {
		for _, v := range validReusability {
			if err := r.Reusability(v); err != nil {
				t.Errorf("%s", err)
			}
			if r.reusability != v {
				t.Errorf("r.Reusability(%s): got %s, want %s", v, r.reusability, v)
			}
		}

		for _, v := range []string{"abc", "0123", "-15", "opent"} {
			if err := r.Reusability(v); err == nil {
				t.Errorf("err shouldn't be nil with r.Reusability(%s)", v)
			}

			if r.reusability == v {
				t.Errorf("r.Reusability(%s): got %s, want %s", v, r.reusability, "")
			}
		}
	})

	t.Run("Change Profile", func(t *testing.T) {
		for _, v := range validProfile {
			if err := r.Profile(v); err != nil {
				t.Errorf("%s", err)
			}

			if r.profile != v {
				t.Errorf("r.Profile(%s): got %s, want %s", v, r.profile, v)
			}
		}

		for _, v := range []string{"abc", "0123", "-15", "standart"} {
			if err := r.Profile(v); err == nil {
				t.Errorf("err shouldn't be nil with r.Profile(%s)", v)
			}

			if r.profile == v {
				t.Errorf("r.Profile(%s): got %s, want %s", v, r.profile, "")
			}
		}
	})

	t.Run("Change Rows", func(t *testing.T) {
		for _, v := range validRows {
			if err := r.Rows(v); err != nil {
				t.Errorf("%s", err)
			}
			if r.rows != v {
				t.Errorf("r.Rows(%s): got %s, want %s", v, r.rows, "")
			}
		}

		for _, v := range []string{"-20", "3.14", "test", "-1"} {
			if err := r.Rows(v); err == nil {
				t.Errorf("err shouldn't be nil with r.Rows(%s)", v)
			}

			if r.rows == v {
				t.Errorf("r.Rows(%s): got %s, want %s", v, r.rows, "")
			}
		}
	})

	t.Run("Change Start", func(t *testing.T) {
		for _, v := range validStart {
			if err := r.Start(v); err != nil {
				t.Errorf("%s", err)
			}

			if r.start != v {
				t.Errorf("r.Start(%s): got %s, want %s", v, r.start, v)
			}
		}

		for _, v := range []string{"-20", "3.14", "test", "0"} {
			if err := r.Start(v); err == nil {
				t.Errorf("err shouldn't be nil with r.Start(%s)", v)
			}

			if r.start == v {
				t.Errorf("r.Start(%s): got %s, want %s", v, r.start, "")
			}
		}
	})
}

func TestFieldChangeCursorSearchRequest(t *testing.T) {
	c := NewClient("abc", "")
	r, _ := NewCursorSearchRequest(c, "", "", "")

	t.Run("Change Reusability", func(t *testing.T) {
		for _, v := range validReusability {
			if err := r.Reusability(v); err != nil {
				t.Error(err)
			}
		}

		for _, v := range []string{"abc", "0123", "-15", "opent"} {
			if err := r.Reusability(v); err == nil {
				t.Errorf("err shouldn't be nil with r.Reusability(%s)", v)
			}

			if r.reusability == v {
				t.Errorf("r.Reusability(%s): got %s, want %s", v, r.reusability, "")
			}
		}
	})

	t.Run("Change Profile", func(t *testing.T) {
		for _, v := range validProfile {
			if err := r.Profile(v); err != nil {
				t.Error(err)
			}
		}

		for _, v := range []string{"abc", "0123", "-15", "standart"} {
			if err := r.Profile(v); err == nil {
				t.Errorf("err shouldn't be nil with r.Profile(%s)", v)
			}

			if r.profile == v {
				t.Errorf("r.Profile(%s): got %s, want %s", v, r.profile, "")
			}
		}
	})

	t.Run("Change Cursor", func(t *testing.T) {
		for i, v := range validCursor {
			r.Cursor(v)
			if i == 0 {
				if r.cursor != "*" {
					t.Errorf("r.Cursor(%s): got %s, want %s", v, r.cursor, "*")
				}
			} else {
				if r.cursor != v {
					t.Errorf("r.Cursor(%s): got %s, want %s", v, r.cursor, v)
				}
			}
		}
	})
}
