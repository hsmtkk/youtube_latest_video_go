package apikey_test

import (
	"os"
	"testing"

	"github.com/hsmtkk/youtube_latest_video_go/pkg/apikey"
	"github.com/stretchr/testify/assert"
)

func TestLoadAPIKey(t *testing.T) {
	loader := apikey.New()
	want := "abcd"
	os.Setenv(apikey.YoutubeAPIKey, want)
	got, err := loader.LoadAPIKey()
	assert.Nil(t, err, "should be nil")
	assert.Equal(t, want, got, "should match")
}
