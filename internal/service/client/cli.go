package client

import "golang.org/x/net/context"

type Cli struct {
	gRpcClient *gRpcClient
}

func NewClient(gRpcAddr, gRpcPort string) *Cli {
	return &Cli{gRpcClient: newGRpcClient(gRpcAddr, gRpcPort)}
}

func (cli *Cli) Execute(ctx context.Context, cliArgs []string) {
	keys := cli.findKeys(cliArgs)
	videoIds := cli.findVideoIds(cliArgs)

	if val, ok := keys["async"]; ok {
		if val {
			cli.gRpcClient.DownloadThumbnailsAsync(ctx, videoIds)
		} else {
			cli.gRpcClient.DownloadThumbnails(ctx, videoIds)
		}
	}

}

func (cli *Cli) findKeys(cliArgs []string) map[string]bool {

	return nil
}

func (cli *Cli) findVideoIds(cliArgs []string) []string {
	return nil
}

func 