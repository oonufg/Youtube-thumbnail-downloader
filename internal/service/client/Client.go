package client

import (
	protobuff "YoutubeThumbnailDownloader/internal/service/api/gen"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type gRpcClient struct {
	gRpcAddr  string
	gRpcPort  string
	client    protobuff.YtThumbGRPCClient
	clientCon *grpc.ClientConn
}

func newGRpcClient(gRpcAddr, gRpcPort string) *gRpcClient {
	grpcClient := &gRpcClient{
		gRpcAddr: gRpcAddr,
		gRpcPort: gRpcPort,
	}
	grpcClient.connect()

	return grpcClient
}

func (client *gRpcClient) DownloadThumbnails(ctx context.Context, videoIds []string) {
	client.client.DownloadThumbnails(ctx, &protobuff.DownloadThumbnailsRequest{
		VideoId: videoIds,
	})
}

func (client *gRpcClient) DownloadThumbnailsAsync(ctx context.Context, videoIds []string) {
	_, err := client.client.DownloadThumbnailsAsync(ctx, &protobuff.DownloadThumbnailsRequest{
		VideoId: videoIds,
	})
	if err != nil {
		log.Println(err)
	}
}

func (client *gRpcClient) connect() {
	log.Println("Starting client...")
	con, err := grpc.Dial(fmt.Sprintf("%s:%s", client.gRpcAddr, client.gRpcPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	log.Printf("Try connecting to gRPC server at %s:%s\n", client.gRpcAddr, client.gRpcPort)
	if err != nil {
		log.Printf("Failed connect to gRPC server at %s:%s\n", client.gRpcAddr, client.gRpcPort)
	}
	client.clientCon = con
	client.client = protobuff.NewYtThumbGRPCClient(con)
	log.Println("Client started")

}

func (client *gRpcClient) Close() {
	log.Println("Closing client connection...")
	client.clientCon.Close()
	log.Println("Client connection closed")
}
