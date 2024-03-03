package domain

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

func DownloadThumblenailsAsync(ctx context.Context, videoIds []string) {
	wg := new(sync.WaitGroup)
	fmt.Println(len(videoIds))
	wg.Add(len(videoIds))
	for _, videoId := range videoIds {
		go downloadAsyncWrapper(ctx, wg, videoId)
	}
	wg.Wait()
}

func downloadAsyncWrapper(ctx context.Context, wg *sync.WaitGroup, videoId string) {
	DownloadThumblenail(videoId)
	defer wg.Done()

	select {
	case <-ctx.Done():
		log.Printf("Timeout while downloading | %s\n", videoId)
		return
	default:
		return
	}
}

func DownloadThumblenail(videoId string) error {
	log.Printf("Start downloading thumblenail | %s... \n", videoId)

	currentWorkDir, _ := os.Getwd()
	downloadDir := fmt.Sprintf("%s/YoutubeThumblenails", currentWorkDir)
	filepath := fmt.Sprintf("%s/%s.jpg", downloadDir, videoId)
	_, err := os.Stat(filepath)
	var file *os.File

	if os.IsNotExist(err) {
		file, _ = os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	} else {
		return errors.New("File already exists")
	}

	thumbelnailBytes, _ := GetThumbnail(videoId)
	io.Copy(file, bytes.NewReader(thumbelnailBytes))

	log.Printf("Finished downloading thumblenail | %s\n", videoId)
	return nil
}
