package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
	"time"
)

type testResponse struct {
	Results []struct {
		Task string `json:"task"`
		Done bool   `json:"done"`
	} `json:"results"`
	Date         time.Time `json:"date"`
	TotalResults int       `json:"total_results"`
}

func setupAPI(t *testing.T) (server *httptest.Server, clean func()) {
	t.Helper()

	tmpFile, err := os.CreateTemp("", "todo_test_*.json")
	if err != nil {
		t.Fatal("Error creating temp file:", err)
	}

	if err := os.WriteFile(tmpFile.Name(), []byte("[]"), 0644); err != nil {
		t.Fatal("Error initializing temp file:", err)
	}

	server = httptest.NewServer(newMux(tmpFile.Name()))

	for i := 1; i <= 3; i++ {
		var body bytes.Buffer
		task := struct {
			Task string `json:"task"`
		}{
			Task: "Test Task " + strconv.Itoa(i),
		}

		if err := json.NewEncoder(&body).Encode(task); err != nil {
			t.Fatalf("Error encoding task %d: %v", i, err)
		}

		resp, err := http.Post(server.URL+"/todo", "application/json", &body)
		if err != nil {
			t.Fatalf("Error posting task %d: %v", i, err)
		}

		if resp.Body != nil {
			resp.Body.Close()
		}

		if resp.StatusCode != http.StatusCreated {
			var errorBody bytes.Buffer
			if resp.Body != nil {
				errorBody.ReadFrom(resp.Body)
			}
			t.Fatalf("Expected status 201 for task %d, got %d. Error: %s", i, resp.StatusCode, errorBody.String())
		}
	}

	clean = func() {
		server.Close()
		os.Remove(tmpFile.Name())
	}

	return server, clean
}

func TestGet(t *testing.T) {
	server, clean := setupAPI(t)
	defer clean()

	testCases := []struct {
		name         string
		path         string
		expItems     int
		expFirstTask string
	}{
		{
			name:         "GetAllTasks",
			path:         "/todo",
			expItems:     3,
			expFirstTask: "Test Task 1",
		},
		{
			name:         "GetSingleTask",
			path:         "/todo/0",
			expItems:     1,
			expFirstTask: "Test Task 1",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := http.Get(server.URL + tc.path)
			if err != nil {
				t.Fatalf("GET %s error: %v", tc.path, err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				t.Fatalf("Expected status 200 for %s, got %d", tc.path, resp.StatusCode)
			}

			contentType := resp.Header.Get("Content-Type")
			if contentType != "application/json" {
				t.Fatalf("Expected application/json, got %s", contentType)
			}

			var response testResponse
			if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
				t.Fatalf("Error decoding response: %v", err)
			}

			if len(response.Results) != tc.expItems {
				t.Errorf("Expected %d items, got %d", tc.expItems, len(response.Results))
			}

			if len(response.Results) > 0 && response.Results[0].Task != tc.expFirstTask {
				t.Errorf("Expected first task '%s', got '%s'", tc.expFirstTask, response.Results[0].Task)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	server, clean := setupAPI(t)
	defer clean()

	t.Run("AddNewTask", func(t *testing.T) {
		var body bytes.Buffer
		newTask := struct {
			Task string `json:"task"`
		}{
			Task: "New Test Task",
		}

		if err := json.NewEncoder(&body).Encode(newTask); err != nil {
			t.Fatalf("Error encoding new task: %v", err)
		}

		resp, err := http.Post(server.URL+"/todo", "application/json", &body)
		if err != nil {
			t.Fatalf("POST error: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusCreated {
			t.Errorf("Expected 201 Created, got %d", resp.StatusCode)
		}
	})

	t.Run("VerifyAddedTask", func(t *testing.T) {
		resp, err := http.Get(server.URL + "/todo/3")
		if err != nil {
			t.Fatalf("GET error: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Expected 200 OK, got %d", resp.StatusCode)
		}

		var response testResponse
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			t.Fatalf("Error decoding response: %v", err)
		}

		if len(response.Results) != 1 {
			t.Fatalf("Expected 1 result, got %d", len(response.Results))
		}

		expectedTask := "New Test Task"
		if response.Results[0].Task != expectedTask {
			t.Errorf("Expected '%s', got '%s'", expectedTask, response.Results[0].Task)
		}
	})
}

func TestDelete(t *testing.T) {
	server, clean := setupAPI(t)
	defer clean()

	t.Run("DeleteFirstTask", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodDelete, server.URL+"/todo/0", nil)
		if err != nil {
			t.Fatalf("Error creating DELETE request: %v", err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("Error executing DELETE request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected 200 OK, got %d", resp.StatusCode)
		}
	})

	t.Run("VerifyDeletion", func(t *testing.T) {
		resp, err := http.Get(server.URL + "/todo")
		if err != nil {
			t.Fatalf("GET error: %v", err)
		}
		defer resp.Body.Close()

		var response testResponse
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			t.Fatalf("Error decoding response: %v", err)
		}

		if len(response.Results) != 2 {
			t.Errorf("Expected 2 results after deletion, got %d", len(response.Results))
		}

		expectedFirstTask := "Test Task 2"
		if len(response.Results) > 0 && response.Results[0].Task != expectedFirstTask {
			t.Errorf("Expected first task '%s', got '%s'", expectedFirstTask, response.Results[0].Task)
		}
	})
}

func TestComplete(t *testing.T) {
	server, clean := setupAPI(t)
	defer clean()

	t.Run("CompleteFirstTask", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPatch, server.URL+"/todo/0", nil)
		if err != nil {
			t.Fatalf("Error creating PATCH request: %v", err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("Error executing PATCH request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected 200 OK, got %d", resp.StatusCode)
		}
	})

	t.Run("VerifyCompletion", func(t *testing.T) {
		resp, err := http.Get(server.URL + "/todo")
		if err != nil {
			t.Fatalf("GET error: %v", err)
		}
		defer resp.Body.Close()

		var response testResponse
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			t.Fatalf("Error decoding response: %v", err)
		}

		if response.TotalResults != 3 {
			t.Errorf("Expected 3 total results, got %d", response.TotalResults)
		}

		if len(response.Results) > 0 && !response.Results[0].Done {
			t.Error("Expected first task to be completed")
		}

		if len(response.Results) > 1 && response.Results[1].Done {
			t.Error("Expected second task to be incomplete")
		}
		if len(response.Results) > 2 && response.Results[2].Done {
			t.Error("Expected third task to be incomplete")
		}
	})
}
