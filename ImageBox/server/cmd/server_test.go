package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func setupAPI(t *testing.T) (string, func()) {
	t.Helper()
	app := &ServerApplication{}
	ts := httptest.NewServer(app.routes())

	return ts.URL, func() {
		ts.Close()
	}
}

func TestGet(t *testing.T) {
	//TODO: update tests
	testCases := []struct {
		name       string
		path       string
		expCode    int
		expContent string
	}{
		{name: "GetRoot", path: "/", expCode: http.StatusOK, expContent: "There's an API here"},
	}

	url, cleanup := setupAPI(t)
	defer cleanup()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var (
				// resp struct {
				// 	Content string `json:"content"`
				// }
				body []byte
				err  error
			)

			r, err := http.Get(url + tc.path)
			if err != nil {
				t.Error(err)
			}

			defer r.Body.Close()

			if body, err = ioutil.ReadAll(r.Body); err != nil {
				t.Error(err)
			}
			fmt.Println(string(body))
			switch {
			case r.Header.Get("Content-Type") == "application/json":
				// if err = json.NewDecoder(r.Body).Decode(&resp); err != nil {
				// 	t.Error(err)
				// }
				// if resp.Content != tc.expContent {
				// 	t.Errorf("Expected %q, got %q", tc.expContent, resp.Content)
				// }
				if !strings.Contains(string(body), tc.expContent) {
					t.Errorf("Expected %q, got %q ", tc.expContent, string(body))
				}
			}
		})
	}

}
