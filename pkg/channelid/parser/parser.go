package parser

import (
	"encoding/json"
	"fmt"
)

type ResponseParser interface {
	ParseResponse(js string) (string, error)
}

func New() ResponseParser {
	return &responseParserImpl{}
}

type responseParserImpl struct{}

type responseSchema struct {
	Items []item `json:"items"`
}

type item struct {
	Snippet struct {
		ChannelID string `json:"channelId"`
	} `json:"snippet"`
}

func (impl *responseParserImpl) ParseResponse(js string) (string, error) {
	sch := responseSchema{}
	err := json.Unmarshal([]byte(js), &sch)
	if err != nil {
		return "", err
	}
	if len(sch.Items) == 0 {
		return "", fmt.Errorf("failed to get channel ID")
	}
	return sch.Items[0].Snippet.ChannelID, nil
}
