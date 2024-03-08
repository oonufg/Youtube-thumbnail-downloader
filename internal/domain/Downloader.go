package domain

import (
	"YoutubeThumbnailDownloader/internal/cache"
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

type ThumbnailDownloader struct {
	cache       cache.ThumbnailCache
	downloadDir string
}

func (thumbnailService *ThumbnailDownloader) DownloadThumbnailsAsync(ctx context.Context, videoIds []string) {
	wg := new(sync.WaitGroup)
	wg.Add(len(videoIds))

	for _, videoId := range videoIds {
		go thumbnailService.downloadThumbnailAsyncWrapper(ctx, wg, videoId)
	}
	wg.Wait()
}

func New(downloadDir string, cache cache.ThumbnailCache) *ThumbnailDownloader {
	return &ThumbnailDownloader{
		downloadDir: downloadDir,
		cache:       cache,
	}
}

func (thumbnailService *ThumbnailDownloader) DownloadThumbnails(ctx context.Context, videoIds []string) {
	for _, videoId := range videoIds {
		thumbnailService.downloadThumbnail(videoId)
	}
}

func (thumbnailService *ThumbnailDownloader) downloadThumbnailAsyncWrapper(ctx context.Context, wg *sync.WaitGroup, videoId string) {
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

func (thumbnailService *ThumbnailDownloader) downloadThumbnail(videoId string) error {
	log.Printf("Start downloading thumbnails | %s... \n", videoId)

	filepath := fmt.Sprintf("%s/%s.jpg", thumbnailService.downloadDir, videoId)
	_, err := os.Stat(filepath)
	var file *os.File

	if os.IsNotExist(err) {
		file, _ = os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	} else {
		log.Printf("Thumbnail already been downloaded | %s\n", videoId)
		return nil
	}

	var thumbnailsBytes []byte
	if thumbnailService.cache.IsThumbnailCached(videoId) {
		log.Printf("Getting thumbnail from cache | %s\n", videoId)
		thumbnailsBytes = thumbnailService.cache.GetThumbnail(context.TODO(), videoId)
	} else {
		thumbnailsBytes, _ = GetThumbnail(videoId)
		thumbnailService.cache.CacheThumbnail(context.TODO(), videoId, thumbnailsBytes)
	}

	io.Copy(file, bytes.NewReader(thumbnailsBytes))
	log.Printf("Finished downloading thumbnails | %s\n", videoId)
	return nil
}
