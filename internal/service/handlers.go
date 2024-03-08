package server

import (
	"YoutubeThumbnailDownloader/internal/domain"
	protobuff "YoutubeThumbnailDownloader/internal/service/api/gen"
	"context"
)

type YtThumbHandler struct {
	protobuff.UnimplementedYtThumbGRPCServer
	thumbDownloader domain.ThumbnailDownloader
}

func MakeYtThumbHandler(thumbDownload *domain.ThumbnailDownloader) *YtThumbHandler {
	return &YtThumbHandler{
		thumbDownloader: *thumbDownload,
	}
}

func (ytThumbHandler *YtThumbHandler) DownloadThumbnails(ctx context.Context, request *protobuff.DownloadThumbnailsRequest) (*protobuff.Empty, error) {
	ytThumbHandler.thumbDownloader.DownloadThumbnails(ctx, request.GetVideoId())
	return &protobuff.Empty{}, nil
}

func (ytThumbHandler *YtThumbHandler) DownloadThumbnailsAsync(ctx context.Context, request *protobuff.DownloadThumbnailsRequest) (*protobuff.Empty, error) {
	ytThumbHandler.thumbDownloader.DownloadThumbnailsAsync(ctx, request.GetVideoId())
	return &protobuff.Empty{}, nil
}
