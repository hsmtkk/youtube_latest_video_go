package geturl

import (
	"log"

	"github.com/hsmtkk/youtube_latest_video_go/pkg/geturl"
	"github.com/spf13/cobra"
)

var GetURLCommand = &cobra.Command{
	Use:  "geturl",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		err := geturl.New().GetURL(url)
		if err != nil {
			log.Fatal(err)
		}
	},
}
