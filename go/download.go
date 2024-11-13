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
	file, err := os.Open("../config/settings.json")
	if err != nil {
		fmt.Println(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		fmt.Println(err)
	}
	return config
}

func DownloadVideo(url, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("Video downloaded to %s\n", filepath)
	return nil
}
