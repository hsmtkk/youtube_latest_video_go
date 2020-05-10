package latestvideo_test

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"os"
	"testing"

	"github.com/hsmtkk/youtube_latest_video_go/pkg"
	"github.com/hsmtkk/youtube_latest_video_go/pkg/apikey"
	"github.com/hsmtkk/youtube_latest_video_go/pkg/latestvideo"
	"github.com/stretchr/testify/assert"
)

func TestLocal(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bs, err := httputil.DumpRequest(r, false)
		assert.Nil(t, err, "should be nil")
		log.Println(string(bs))
		reader, err := os.Open("./example_response.json")
		defer reader.Close()
		assert.Nil(t, err, "should be nil")
		written, err := io.Copy(w, reader)
		assert.Nil(t, err, "should be nil")
		assert.Greater(t, written, int64(0), "should be greater than zero")
	}))
	defer ts.Close()

	getter := latestvideo.NewForTest(ts.Client(), ts.URL, "apiKey")
	want := pkg.Video{VideoID: "iuB8oL1ytTk", ThumbnailURL: "https://i.ytimg.com/vi/iuB8oL1ytTk/default.jpg"}
	got, err := getter.GetLatestVideo("channelID")
	assert.Nil(t, err, "should be nil")
	assert.Equal(t, want, got, "should be equal")
}

func TestReal(t *testing.T) {
	apiKey, err := apikey.New().LoadAPIKey()
	assert.Nil(t, err, "should be nil")
	channelID := "UCK8sQmJBp8GCxrOtXWBpyEA"
	getter := latestvideo.New(apiKey)
	video, err := getter.GetLatestVideo(channelID)
	assert.Nil(t, err, "should be nil")
	assert.NotEmpty(t, video.VideoID, "should not be empty")
	assert.NotEmpty(t, video.ThumbnailURL, "should not be empty")
}
