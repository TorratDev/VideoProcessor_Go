# Video Processor

This project downloads a video from a specified URL and converts it to audio.

## Requirements
- Go 1.18+
- ffmpeg (for video to audio conversion)

## Setup

1. Install Go and `ffmpeg`.
2. Configure `config/settings.json` with your desired download paths and video URL.
3. Run the service with:

```bash
   go run main.go
```
