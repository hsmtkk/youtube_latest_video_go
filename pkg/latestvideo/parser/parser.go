package parser

import (
	"encoding/json"
	"fmt"

	"github.com/hsmtkk/youtube_latest_video_go/pkg"
)

type ResponseParser interface {
	ParseResponse(js string) (pkg.Video, error)
}

func New() ResponseParser {
	return &responseParserImpl{}
}

type responseParserImpl struct{}

type responseSchema struct {
	Items []item `json:"items"`
}

type item struct {
	ID struct {
		VideoID string `json:"videoId"`
	} `json:"id"`
	Snippet struct {
		Thumbnails struct {
			Default struct {
				URL string `json:"url"`
			} `json:"default"`
		} `json:"thumbnails"`
	} `json:"snippet"`
}

func (impl *responseParserImpl) ParseResponse(js string) (pkg.Video, error) {
	sch := responseSchema{}
	err := json.Unmarshal([]byte(js), &sch)
	if err != nil {
		return pkg.Video{}, err
	}
	if len(sch.Items) == 0 {
		return pkg.Video{}, fmt.Errorf("failed to get video")
	}
	item := sch.Items[0]
	return pkg.Video{VideoID: item.ID.VideoID, ThumbnailURL: item.Snippet.Thumbnails.Default.URL}, nil
}
