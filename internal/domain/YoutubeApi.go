package domain

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetThumbnail(videoId string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("https://img.youtube.com/vi/%s/hqdefault.jpg", videoId))
	if err != nil {
		log.Println("Error while downloading video with id - ", videoId)
		return nil, fmt.Errorf("Error while downloading video | Video id - %s", videoId)
	}
	defer resp.Body.Close()

	thumblenailBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading bytes from video with id - ", videoId)
		return nil, fmt.Errorf("Error while read response | Video id - %s", videoId)
	}

	return thumblenailBytes, nil
}
