package main

import (
	"YoutubeThumbnailDownloader/internal/cache"
	_ "YoutubeThumbnailDownloader/internal/cache"
	"YoutubeThumbnailDownloader/internal/domain"
	_ "YoutubeThumbnailDownloader/internal/domain"
	"YoutubeThumbnailDownloader/internal/service/client"
	server "YoutubeThumbnailDownloader/internal/service/server"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	sigChan := getListenOsSigChan()
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go RunServer(ctx, wg)
	time.Sleep(1 * time.Second)

	cli := client.NewClient("127.0.0.1", "8080")
	go cli.Execute(ctx, os.Args[1:])

	<-sigChan
	cancel()
	wg.Wait()
}

func getListenOsSigChan() chan os.Signal {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	return sigChan
}

func RunServer(ctx context.Context, wg *sync.WaitGroup) {
	cache, _ := cache.NewCache("cache.sqlite")
	downloader := domain.New("YoutubeThumbnails", cache)
	handler := server.MakeYtThumbHandler(downloader)
	server := server.MakeServer(handler, "127.0.0.1", "8080")

	go server.Run(ctx)
	<-ctx.Done()
	server.Shutdown()
	wg.Done()
}

func CreateAllNecessaryIfNotExists() {
	currentWorkDir, _ := os.Getwd()
	downloadDir := fmt.Sprintf("%s/YoutubeThumbnails", currentWorkDir)
	if _, err := os.Stat(downloadDir); err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir("YoutubeThumbnails", 0777)
			if err != nil {
				log.Fatalln("Can't create YoutubeThumbnails folder")
			}
		}
	}
}

func GetDirToDownloadThumbnails() string {
	currentWorkDir, _ := os.Getwd()
	downloadDir := fmt.Sprintf("%s/YoutubeThumbnails", currentWorkDir)
	return downloadDir
}
