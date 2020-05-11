package atonce

import (
	"github.com/hsmtkk/youtube_latest_video_go/cmd/youtube/channelid"
	"github.com/hsmtkk/youtube_latest_video_go/cmd/youtube/geturl"
	"github.com/hsmtkk/youtube_latest_video_go/cmd/youtube/latestvideo"
	"github.com/spf13/cobra"
)

var AtOnceCommand = &cobra.Command{
	Use:  "atonce",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		channelName := args[0]
		channelID := channelid.GetChannelID(channelName)
		_, videoURL := latestvideo.GetLatestVideo(channelID)
		geturl.GetURL(videoURL)
	},
}
