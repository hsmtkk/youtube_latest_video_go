package channelid

import (
	"io/ioutil"
	"net/http"

	"github.com/hsmtkk/youtube_latest_video_go/pkg/channelid/parser"
)

type ChannelIDResolver interface {
	ResolveChannelID(channelName string) (string, error)
}

func New(apiKey string) ChannelIDResolver {
	url := "https://www.googleapis.com/youtube/v3/search"
	return &channelIDResolverImpl{client: http.DefaultClient, url: url, apiKey: apiKey}
}

func NewForTest(client *http.Client, url, apiKey string) ChannelIDResolver {
	return &channelIDResolverImpl{client: client, url: url, apiKey: apiKey}
}

type channelIDResolverImpl struct {
	client *http.Client
	url    string
	apiKey string
}

func (impl *channelIDResolverImpl) ResolveChannelID(channelName string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, impl.url, nil)
	if err != nil {
		return "", err
	}
	query := req.URL.Query()
	query.Add("part", "id,snippet")
	query.Add("key", impl.apiKey)
	query.Add("type", "channel")
	query.Add("q", channelName)
	req.URL.RawQuery = query.Encode()
	resp, err := impl.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	id, err := parser.New().ParseResponse(string(bs))
	if err != nil {
		return "", err
	}
	return id, nil
}
