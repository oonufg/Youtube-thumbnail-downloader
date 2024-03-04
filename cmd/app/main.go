package main

import (
	"YoutubeThumbnailDownloader/internal/domain"
	"YoutubeThumbnailDownloader/internal/persistence"
	"context"
	"fmt"
	"os"
)

func main() {
	cache, _ := persistence.New("cache.sqlite")
	videos := [1]string{"wC9nSR73_HU"}
	currentWorkDir, _ := os.Getwd()
	downloadDir := fmt.Sprintf("%s/YoutubeThumbnails", currentWorkDir)

	service := domain.New(downloadDir, cache)
	service.DownloadThumbnails(context.Background(), videos[:])

}
