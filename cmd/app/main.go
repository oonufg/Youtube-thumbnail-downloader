package main

import (
	_ "YoutubeThumbnailDownloader/internal/cache"
	_ "YoutubeThumbnailDownloader/internal/domain"
	"fmt"
	"log"
	"os"
)

func main() {
	CreateAllNeededIfNotExists()

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
