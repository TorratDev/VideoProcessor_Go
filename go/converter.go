package main

import (
	"fmt"
	"os/exec"
)

func ConvertVideoToAudio(inputPath string, outputPath string) error {
	// Using ffmpeg for conversion; make sure it's installed on the system
	cmd := exec.Command("ffmpeg", "-i", inputPath, outputPath)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to convert video to audio: %v", err)
	}

	fmt.Printf("Audio saved to %s\n", outputPath)
	return nil
}
