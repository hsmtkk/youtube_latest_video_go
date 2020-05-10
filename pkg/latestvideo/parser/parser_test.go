package parser_test

import (
	"io/ioutil"
	"testing"

	"github.com/hsmtkk/youtube_latest_video_go/pkg"
	"github.com/hsmtkk/youtube_latest_video_go/pkg/latestvideo/parser"
	"github.com/stretchr/testify/assert"
)

func TestParseResponse(t *testing.T) {
	js, err := ioutil.ReadFile("../example_response.json")
	assert.Nil(t, err, "should be nil")
	want := pkg.Video{VideoID: "iuB8oL1ytTk", ThumbnailURL: "https://i.ytimg.com/vi/iuB8oL1ytTk/default.jpg"}
	got, err := parser.New().ParseResponse(string(js))
	assert.Nil(t, err, "should be nil")
	assert.Equal(t, want, got, "should be equal")
}
