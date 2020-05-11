package latestvideo

import (
	"io/ioutil"
	"net/http"

	"github.com/hsmtkk/youtube_latest_video_go/pkg"
	"github.com/hsmtkk/youtube_latest_video_go/pkg/latestvideo/parser"
)

type LatestVideoGetter interface {
	GetLatestVideo(channelID string) (pkg.Video, error)
}

func New(apiKey string) LatestVideoGetter {
	return &latestVideoGetterImpl{client: http.DefaultClient, url: pkg.SearchAPIURL, apiKey: apiKey}
}

func NewForTest(client *http.Client, url string, apiKey string) LatestVideoGetter {
	return &latestVideoGetterImpl{client: client, url: url, apiKey: apiKey}
}

type latestVideoGetterImpl struct {
	client *http.Client
	url    string
	apiKey string
}

func (impl *latestVideoGetterImpl) GetLatestVideo(channelID string) (pkg.Video, error) {
	req, err := http.NewRequest(http.MethodGet, impl.url, nil)
	if err != nil {
		return pkg.Video{}, err
	}
	query := req.URL.Query()
	query.Add("part", "id,snippet")
	query.Add("key", impl.apiKey)
	query.Add("channelId", channelID)
	query.Add("maxResults", "1")
	query.Add("order", "date")
	req.URL.RawQuery = query.Encode()
	resp, err := impl.client.Do(req)
	if err != nil {
		return pkg.Video{}, err
	}
	defer resp.Body.Close()
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return pkg.Video{}, err
	}
	video, err := parser.New().ParseResponse(string(bs))
	if err != nil {
		return pkg.Video{}, err
	}
	return video, nil
}
