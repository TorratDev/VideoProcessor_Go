package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestConvertVideoToAudio(t *testing.T) {
	// Create a temporary directory for input/output files
	tempDir := t.TempDir()
	inputPath := filepath.Join("../assets/test_video.mp4")
	outputPath := filepath.Join(tempDir, "output.mp3")

	// Run the ConvertVideoToAudio function
	err := ConvertVideoToAudio(inputPath, outputPath)
	if err != nil {
		t.Fatalf("ConvertVideoToAudio failed: %v", err)
	}

	// Check if the audio file was created
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		t.Errorf("Expected output file was not created")
	}
}
