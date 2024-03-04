package main

import (
	"YoutubeThumbnailDownloader/internal/domain"
	"YoutubeThumbnailDownloader/internal/persistence"
	"context"
	"fmt"
	"log"
	"os"
)

func main() {
	CreateAllNeededIfNotExists()
	cache, _ := persistence.NewCache("cache.sqlite")
	videos := [1]string{"wC9nSR73_HU"}
	currentWorkDir, _ := os.Getwd()
	downloadDir := fmt.Sprintf("%s/YoutubeThumbnails", currentWorkDir)

	service := domain.New(downloadDir, cache)
	service.DownloadThumbnails(context.Background(), videos[:])

}

func CreateAllNeededIfNotExists() {
	currentWorkDir, _ := os.Getwd()
	downloadDir := fmt.Sprintf("%s/YoutubeThumbnails", currentWorkDir)
	if _, err := os.Stat(downloadDir); err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir("YoutubeThumbnails", 0777)
			if err != nil {
				log.Fatalln("YoutubeThumbnails folder not created")
			}
		}
	}
}
