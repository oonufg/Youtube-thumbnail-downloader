package server

import (
	protobuff "YoutubeThumbnailDownloader/internal/service/api/gen"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	GRPCAdder      string
	GRPCPort       string
	ytThumbHandler *YtThumbHandler
	gRpcServer     *grpc.Server
}

func MakeServer(ytThumbHandler *YtThumbHandler) *Server {
	return &Server{
		ytThumbHandler: ytThumbHandler,
	}
}

func (server *Server) Run(ctx context.Context) {
	log.Println("Run gRPC server..")
	gRpcServer := grpc.NewServer()
	protobuff.RegisterYtThumbGRPCServer(gRpcServer, server.ytThumbHandler)
	gRpcFullAddr := fmt.Sprintf("%s:%s", server.GRPCAdder, server.GRPCPort)

	listener, err := net.Listen("tcp", gRpcFullAddr)
	if err != nil {
		log.Fatalf("Failed to start gRPC at %s\n", gRpcFullAddr)
	}

	gRpcServer.Serve(listener)
	server.gRpcServer = gRpcServer
}

func (server *Server) Shutdown() {
	log.Println("Shooting down gRPC server..")
	server.gRpcServer.GracefulStop()
}
