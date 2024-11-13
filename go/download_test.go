package main

import (
	"bytes"
	"io"
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

	file, err := os.Open("../assets/test_video.mp4")
	if err != nil {
		t.Fatal(err)
	}

	expectedData, err := io.ReadAll(file)
	if err != nil {
		t.Fatal(err)
	}
	// Set up a test server to simulate downloading a video
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write(expectedData)
		if err != nil {
			return
		}
	}))
	defer server.Close()

	// Run the DownloadVideo function
	err = DownloadVideo(server.URL, tempFile)
	if err != nil {
		t.Fatalf("DownloadVideo failed: %v", err)
	}

	// Check if the file exists and contains the correct data
	data, err := os.ReadFile(tempFile)
	if err != nil {
		t.Fatalf("Failed to read downloaded file: %v", err)
	}

	if !bytes.Equal(data, expectedData) {
		t.Errorf("Downloaded data does not match")
	}
}
