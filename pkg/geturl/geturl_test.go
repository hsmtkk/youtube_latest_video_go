package geturl_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/hsmtkk/youtube_latest_video_go/pkg/geturl"
	"github.com/stretchr/testify/assert"
)

func TestReal(t *testing.T) {
	getter := geturl.New()
	err := getter.GetURL("https://i.ytimg.com/vi/iuB8oL1ytTk/default.jpg")
	assert.Nil(t, err, "should be nil")
	err = os.Remove("default.jpg")
	assert.Nil(t, err, "should be nil")
}

func TestLocal(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "test")
	}))
	defer ts.Close()

	getter := geturl.NewForTest(ts.Client())
	err := getter.GetURLTo(ts.URL, "test")
	assert.Nil(t, err, "should be nil")
	err = os.Remove("test")
	assert.Nil(t, err, "should be nil")
}
