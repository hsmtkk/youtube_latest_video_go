package latestvideo

import (
	"fmt"
	"log"

	"github.com/hsmtkk/youtube_latest_video_go/pkg/apikey"
	"github.com/hsmtkk/youtube_latest_video_go/pkg/latestvideo"
	"github.com/spf13/cobra"
)

var LatestVideoCommand = &cobra.Command{
	Use:  "latestvideo",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		channelID := args[0]
		apiKey, err := apikey.New().LoadAPIKey()
		if err != nil {
			log.Fatal(err)
		}
		getter := latestvideo.New(apiKey)
		video, err := getter.GetLatestVideo(channelID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Video ID: %s\n", video.VideoID)
		fmt.Printf("Thumbnail URL: %s\n", video.ThumbnailURL)
	},
}
