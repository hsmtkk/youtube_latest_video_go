package main

import (
	"log"

	"github.com/hsmtkk/youtube_latest_video_go/cmd/youtube/channelid"
	"github.com/hsmtkk/youtube_latest_video_go/cmd/youtube/geturl"
	"github.com/hsmtkk/youtube_latest_video_go/cmd/youtube/latestvideo"
	"github.com/spf13/cobra"
)

func main() {
	rootCommand := &cobra.Command{
		Use: "youtube",
	}
	rootCommand.AddCommand(channelid.ChannelIDCommand)
	rootCommand.AddCommand(geturl.GetURLCommand)
	rootCommand.AddCommand(latestvideo.LatestVideoCommand)
	if err := rootCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}
