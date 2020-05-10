package parser_test

import (
	"io/ioutil"
	"testing"

	"github.com/hsmtkk/youtube_latest_video_go/pkg/channelid/parser"
	"github.com/stretchr/testify/assert"
)

func TestParseResponse(t *testing.T) {
	js, err := ioutil.ReadFile("../example_response.json")
	assert.Nil(t, err, "should be nil")
	want := "UCK8sQmJBp8GCxrOtXWBpyEA"
	got, err := parser.New().ParseResponse(string(js))
	assert.Nil(t, err, "should be nil")
	assert.Equal(t, want, got, "should be equal")
}
