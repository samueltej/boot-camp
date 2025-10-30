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

func setupAPI(t *testing.T) (url string, cleanup func()) {
	t.Helper()

	server := httptest.NewServer(newMux("/todo"))
	url = server.URL

	cleanup = func() { server.Close() }

	return url, cleanup
}

func TestGet(t *testing.T) {
	cases := []testCase{
		{
			name:            "GET Root",
			path:            "/",
			expectedCode:    http.StatusOK,
			expectedContent: "Hello World!!",
		},
		{
			name:            "GET Not Found",
			path:            "/no-exists",
			expectedCode:    http.StatusNotFound,
			expectedContent: "404",
		},
	}

	url, cleanup := setupAPI(t)
	defer cleanup()

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := http.Get(url + tc.path)
			if err != nil {
				t.Fatalf("GET request failed: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tc.expectedCode {
				t.Errorf(
					"expected %d (%s), got %d (%s)",
					tc.expectedCode, http.StatusText(tc.expectedCode),
					resp.StatusCode, http.StatusText(resp.StatusCode),
				)
			}

			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("reading body failed: %v", err)
			}
			body := string(bodyBytes)

			switch resp.Header.Get("Content-Type") {
			case "text/plain; charset=utf-8":
				if !strings.Contains(body, tc.expectedContent) {
					t.Errorf("expected content %q, got %q", tc.expectedContent, body)
				}
			default:
				t.Fatalf("unsupported Content-Type: %q", resp.Header.Get("Content-Type"))
			}
		})
	}
}
