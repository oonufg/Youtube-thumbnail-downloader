package client

import (
	"regexp"
	"strings"

	"golang.org/x/net/context"
)

type Cli struct {
	gRpcClient       *gRpcClient
	findVideoIdRegEx *regexp.Regexp
	findKeysRegEx    *regexp.Regexp
}

func NewClient(gRpcAddr, gRpcPort string) *Cli {
	findVideoIdRegEx := regexp.MustCompile(`https://www.youtube.com/watch\?v=[\s\S]*`)
	findKeysRegEx := regexp.MustCompile(`--[a-zA-Z]*`)
	return &Cli{
		gRpcClient:       newGRpcClient(gRpcAddr, gRpcPort),
		findVideoIdRegEx: findVideoIdRegEx,
		findKeysRegEx:    findKeysRegEx,
	}
}

func (cli *Cli) Execute(ctx context.Context, cliArgs []string) {
	keys := cli.findKeys(cliArgs)
	videoIds := cli.findVideoIds(cliArgs)

	if _, ok := keys["async"]; ok {
		cli.gRpcClient.DownloadThumbnailsAsync(ctx, videoIds)
	} else {
		cli.gRpcClient.DownloadThumbnails(ctx, videoIds)
	}
}

func (cli *Cli) findKeys(cliArgs []string) map[string]struct{} {
	keys := make(map[string]struct{})
	for _, val := range cliArgs {
		if cli.findKeysRegEx.MatchString(val) {
			key := strings.Split(val, "--")[1]
			keys[key] = struct{}{}
		}
	}
	return keys
}

func (cli *Cli) findVideoIds(cliArgs []string) []string {
	videoIds := make([]string, 0)
	for _, val := range cliArgs {
		if cli.findVideoIdRegEx.MatchString(val) {
			videoId := strings.Split(val, "watch?v=")[1]
			if videoId != "" {
				videoIds = append(videoIds, videoId)
			}
		}
	}
	return videoIds
}
