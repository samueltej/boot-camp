package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func setupAPI(t *testing.T) (baseURL string, cleanup func()) {
	t.Helper()
	testServer := httptest.NewServer(newMux())
	return testServer.URL, func() {
		testServer.Close()
	}
}

func TestGetRoot(t *testing.T) {
	targetPath := "/"
	expectedStatus := http.StatusOK
	expectedText := "Hello World!!"
	
	serverURL, closeServer := setupAPI(t)
	defer closeServer()
	
	response, err := http.Get(serverURL + targetPath)
	if err != nil {
		t.Fatalf("Failed to execute GET request: %v", err)
	}
	defer response.Body.Close()
	
	if response.StatusCode != expectedStatus {
		t.Errorf("Expected HTTP status %d (%s), but got %d (%s)", 
			expectedStatus, http.StatusText(expectedStatus), 
			response.StatusCode, http.StatusText(response.StatusCode))
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	bodyString := string(responseBody)

	if !strings.Contains(bodyString, expectedText) {
		t.Errorf("Expected to find %q in response, but got %q", expectedText, bodyString)
	}
}

func TestGetNotFound(t *testing.T) {
	targetPath := "/undefined-route"
	expectedStatus := http.StatusNotFound
	expectedText := "404"
	
	serverURL, closeServer := setupAPI(t)
	defer closeServer()
	
	response, err := http.Get(serverURL + targetPath)
	if err != nil {
		t.Fatalf("Failed to execute GET request: %v", err)
	}
	defer response.Body.Close()
	
	if response.StatusCode != expectedStatus {
		t.Errorf("Expected HTTP status %d (%s), but got %d (%s)",
			expectedStatus, http.StatusText(expectedStatus),
			response.StatusCode, http.StatusText(response.StatusCode))
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	bodyString := string(responseBody)
	
	if !strings.Contains(bodyString, expectedText) {
		t.Errorf("Expected to find %q in response, but got %q", expectedText, bodyString)
	}
}