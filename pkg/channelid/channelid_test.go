package channelid_test

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"testing"

	"github.com/hsmtkk/youtube_latest_video_go/pkg/apikey"
	"github.com/hsmtkk/youtube_latest_video_go/pkg/channelid"
	"github.com/stretchr/testify/assert"
)

func TestLocal(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bs, err := httputil.DumpRequest(r, false)
		assert.Nil(t, err, "should be nil")
		log.Println(string(bs))
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()

	resolver := channelid.NewForTest(ts.Client(), ts.URL, "dummy")
	channelID, err := resolver.ResolveChannelID("Google")
	assert.Nil(t, err, "should be nil")
	log.Println(channelID)
}

func TestReal(t *testing.T) {
	apiKey, err := apikey.New().LoadAPIKey()
	assert.Nil(t, err, "should be nil")
	resolver := channelid.New(apiKey)
	want := "UCK8sQmJBp8GCxrOtXWBpyEA"
	got, err := resolver.ResolveChannelID("Google")
	assert.Nil(t, err, "should be nil")
	assert.Equal(t, want, got, "should be equal")
}
