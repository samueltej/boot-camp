package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type testCase struct {
	name            string
	path            string
	expectedCode    int
	expectedContent string
}

func setupAPI(t *testing.T) (url string, cleaner func()) {
	t.Helper()

	server := httptest.NewServer(newMux())
	
	url = server.URL

	cleaner = func() {
		server.Close()
	}

	return url, cleaner
}

func TestGet(t *testing.T) {
	cases := []testCase{
		{
			name:            "Root",
			path:            "/",
			expectedCode:    http.StatusOK,
			expectedContent: "Hello World!!",
		},
		{
			name:            "Not Found",
			path:            "/no-exists",
			expectedCode:    http.StatusNotFound,
			expectedContent: "404",
		},
	}

	url, cleaner := setupAPI(t)
	defer cleaner()

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := http.Get(url + tc.path)
			if err != nil {
				t.Fatalf("Error in Get: %v", err)
			}
			defer res.Body.Close()

			if res.StatusCode != tc.expectedCode {
				t.Errorf("expected code %d (%s), but received %d (%s)", 
					tc.expectedCode, http.StatusText(tc.expectedCode), 
					res.StatusCode, http.StatusText(res.StatusCode))
			}

			bodyBytes, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("error reading response body: %v", err)
			}
			body := string(bodyBytes)

			switch res.Header.Get("Content-Type") {
			case "text/plain; charset=utf-8":
				if !strings.Contains(body, tc.expectedContent) {
					t.Errorf("expected content %q, but received %q", tc.expectedContent, body)
				}
			default:
				t.Fatalf("Unsupported Content-Type: %q", res.Header.Get("Content-Type"))
			}
		})
	}
}