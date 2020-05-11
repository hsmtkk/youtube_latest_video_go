package geturl

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
)

type URLGetter interface {
	GetURL(url string) error
	GetURLTo(url, outputPath string) error
}

func New() URLGetter {
	return &urlGetterImpl{client: http.DefaultClient}
}

func NewForTest(client *http.Client) URLGetter {
	return &urlGetterImpl{client: client}
}

type urlGetterImpl struct {
	client *http.Client
}

func (impl *urlGetterImpl) GetURL(u string) error {
	uobj, err := url.Parse(u)
	if err != nil {
		return err
	}
	name := path.Base(uobj.Path)
	return impl.GetURLTo(u, name)
}

func (impl *urlGetterImpl) GetURLTo(url, outputPath string) error {
	resp, err := impl.client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.OpenFile(outputPath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer out.Close()

	written, err := io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	if written <= 0 {
		return fmt.Errorf("IO failure")
	}
	return nil
}
