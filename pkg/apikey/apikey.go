package apikey

import (
	"fmt"
	"os"
)

const YoutubeAPIKey = "YOUTUBE_API_KEY"

type APIKeyLoader interface {
	LoadAPIKey() (string, error)
}

func New() APIKeyLoader {
	return &apiKeyLoaderImpl{}
}

type apiKeyLoaderImpl struct{}

func (imp *apiKeyLoaderImpl) LoadAPIKey() (string, error) {
	val := os.Getenv(YoutubeAPIKey)
	if val == "" {
		return "", fmt.Errorf("environment variable %s is not defined", YoutubeAPIKey)
	}
	return val, nil
}
