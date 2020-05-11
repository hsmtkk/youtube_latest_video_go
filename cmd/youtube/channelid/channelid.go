package channelid

import (
	"fmt"
	"log"

	"github.com/hsmtkk/youtube_latest_video_go/pkg/apikey"
	"github.com/hsmtkk/youtube_latest_video_go/pkg/channelid"
	"github.com/spf13/cobra"
)

var ChannelIDCommand = &cobra.Command{
	Use:  "channelid",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		channelName := args[0]
		fmt.Println(GetChannelID(channelName))
	},
}

func GetChannelID(channelName string) string {
	apiKey, err := apikey.New().LoadAPIKey()
	if err != nil {
		log.Fatal(err)
	}
	resolver := channelid.New(apiKey)
	channelID, err := resolver.ResolveChannelID(channelName)
	if err != nil {
		log.Fatal(err)
	}
	return channelID
}
