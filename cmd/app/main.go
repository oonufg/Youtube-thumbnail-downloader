package main

import (
	"YoutubeThumblenailDownloader/internal/domain"
	"context"
)

func main() {
	videos := [3]string{"aPBKWTNGylk", "-oQQI1bSp-Y", "wC9nSR73_HU"}
	domain.DownloadThumblenailsAsync(context.Background(), videos[:])
}
