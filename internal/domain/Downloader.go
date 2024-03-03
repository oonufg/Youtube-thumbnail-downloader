package domain

import (
	"YoutubeThumbnailDownloader/internal/persistence"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

type ThumbnailService struct {
	cache persistence.ThumbnailCache
}

func (thumbnailService *ThumbnailService) DownloadThumbnailsAsync(ctx context.Context, videoIds []string) {
	wg := new(sync.WaitGroup)
	wg.Add(len(videoIds))

	for _, videoId := range videoIds {
		go thumbnailService.downloadAsyncWrapper(ctx, wg, videoId)
	}
	wg.Wait()
}

func New(cache persistence.ThumbnailCache) *ThumbnailService {
	return &ThumbnailService{cache: cache}
}

func (thumbnailService *ThumbnailService) DownloadThumbnails(ctx context.Context, videoIds []string) {
	for _, videoId := range videoIds {
		thumbnailService.downloadThumbnail(videoId)
	}
}

func (thumbnailService *ThumbnailService) downloadAsyncWrapper(ctx context.Context, wg *sync.WaitGroup, videoId string) {
	thumbnailService.downloadThumbnail(videoId)
	defer wg.Done()

	select {
	case <-ctx.Done():
		log.Printf("Timeout while downloading | %s\n", videoId)
		return
	default:
		return
	}
}

func (thumbnailService *ThumbnailService) downloadThumbnail(videoId string) error {
	log.Printf("Start downloading thumbnails | %s... \n", videoId)

	currentWorkDir, _ := os.Getwd()
	downloadDir := fmt.Sprintf("%s/YoutubeThumbnails", currentWorkDir)
	filepath := fmt.Sprintf("%s/%s.jpg", downloadDir, videoId)
	_, err := os.Stat(filepath)
	var file *os.File

	if os.IsNotExist(err) {
		file, _ = os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	} else {
		return errors.New("File already exists")
	}

	var thumbnailsBytes []byte
	if thumbnailService.cache.IsThumbnailCached(videoId) {
		thumbnailsBytes = thumbnailService.cache.GetThumbnail(context.TODO(), videoId)
	} else {
		thumbnailsBytes, _ = GetThumbnail(videoId)
		thumbnailService.cache.CacheThumbnail(context.TODO(), videoId, thumbnailsBytes)
	}

	io.Copy(file, bytes.NewReader(thumbnailsBytes))
	log.Printf("Finished downloading thumbnails | %s\n", videoId)
	return nil
}
