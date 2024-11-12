package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestDownloadVideo(t *testing.T) {
	// Create a temporary directory for downloading
	tempDir := t.TempDir()
	tempFile := filepath.Join(tempDir, "test_video.mp4")

	// Set up a test server to simulate downloading a video
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("mock video data"))
		if err != nil {
			return
		}
	}))
	defer server.Close()

	// Run the DownloadVideo function
	err := DownloadVideo(server.URL, tempFile)
	if err != nil {
		t.Fatalf("DownloadVideo failed: %v", err)
	}

	// Check if the file exists and contains the correct data
	data, err := os.ReadFile(tempFile)
	if err != nil {
		t.Fatalf("Failed to read downloaded file: %v", err)
	}
	expectedData := "mock video data"
	if string(data) != expectedData {
		t.Errorf("Downloaded data does not match, got %s, want %s", string(data), expectedData)
	}
}
