package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestConvertVideoToAudio(t *testing.T) {
	// Create a temporary directory for input/output files
	tempDir := t.TempDir()
	inputPath := filepath.Join(tempDir, "input.mp4")
	outputPath := filepath.Join(tempDir, "output.mp3")

	// Create an empty video file as a mock input
	file, err := os.Create(inputPath)
	if err != nil {
		t.Fatalf("Failed to create mock video file: %v", err)
	}

	err = file.Close()
	if err != nil {
		return
	}

	// Run the ConvertVideoToAudio function
	err = ConvertVideoToAudio(inputPath, outputPath)
	if err != nil {
		t.Fatalf("ConvertVideoToAudio failed: %v", err)
	}

	// Check if the audio file was created
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		t.Errorf("Expected output file was not created")
	}
}
