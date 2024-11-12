package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// Load config
	config := LoadConfig()

	// Step 1: Download the video
	videoFilePath := filepath.Join(config.DownloadPath, "video.mp4")
	if err := DownloadVideo(config.VideoSourceURL, videoFilePath); err != nil {
		fmt.Printf("Failed to download video: %v\n", err)
		return
	}

	// Step 2: Convert video to audio
	audioFilePath := filepath.Join(config.DownloadPath, "audio.mp3")
	if err := ConvertVideoToAudio(videoFilePath, audioFilePath); err != nil {
		fmt.Printf("Failed to convert video to audio: %v\n", err)
		return
	}

	fmt.Println("Video processing completed successfully!")
}
