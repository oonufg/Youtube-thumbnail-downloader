package main

import (
	"YoutubeThumbnailDownloader/internal/domain"
	"YoutubeThumbnailDownloader/internal/persistence"
	"context"
)

func main() {
	cache, _ := persistence.New("cache.sqlite")
	videos := [1]string{"wC9nSR73_HU"}
	service := domain.New(cache)
	service.DownloadThumbnails(context.Background(), videos[:])

}
