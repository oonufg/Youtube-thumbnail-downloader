package main

import (
	"YoutubeThumbnailDownloader/internal/cache"
	_ "YoutubeThumbnailDownloader/internal/cache"
	"YoutubeThumbnailDownloader/internal/domain"
	_ "YoutubeThumbnailDownloader/internal/domain"
	server "YoutubeThumbnailDownloader/internal/service"
	"context"
	"fmt"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	cache, _ := cache.NewCache("cache.sqlite")
	downloader := domain.New("YoutubeThumbnails", cache)
	handler := server.MakeYtThumbHandler(downloader)
	server := server.MakeServer(handler, "127.0.0.1", "8080")
	server.Run(ctx)
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

func GetDirToDownloadThumbnails() string {
	currentWorkDir, _ := os.Getwd()
	downloadDir := fmt.Sprintf("%s/YoutubeThumbnails", currentWorkDir)
	return downloadDir
}
