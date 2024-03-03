package main

import (
	"YoutubeThumbnailDownloader/internal/domain"
	"YoutubeThumbnailDownloader/internal/persistence"

	"context"
	"log"
)

func main() {
	cache, err := persistence.New("cache.sqlite")
	if err != nil {
		log.Fatalln(err)
	}
	tBytes, _ := domain.GetThumbnail("aPBKWTNGylk")
	cache.CacheThumbnail(context.TODO(), "aPBKWTNGylk", tBytes)

}
