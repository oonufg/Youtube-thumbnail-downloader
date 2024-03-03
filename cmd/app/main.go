package main

import (
	"YoutubeThumbnailDownloader/internal/domain"
	"context"
)

func main() {
	videos := [3]string{"aPBKWTNGylk", "-oQQI1bSp-Y", "wC9nSR73_HU"}
	domain.DownloadThumbnails(context.Background(), videos[:])
}
