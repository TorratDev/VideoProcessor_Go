package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Config struct {
	VideoSourceURL string `json:"videoSourceURL"`
	DownloadPath   string `json:"downloadPath"`
}

func LoadConfig() Config {
	var config Config
	file, _ := os.Open("../config/settings.json")
	defer file.Close()
	json.NewDecoder(file).Decode(&config)
	return config
}

func DownloadVideo(url, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("Video downloaded to %s\n", filepath)
	return nil
}
