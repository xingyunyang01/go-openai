package openai

import (
	"log"
	"net/http"
	"net/url"
)

const (
	apiURLv1                       = "https://api.openai.com/v1"
	defaultEmptyMessagesLimit uint = 300
)

// ClientConfig is a configuration of a client.
type ClientConfig struct {
	authToken string

	HTTPClient *http.Client

	BaseURL string
	OrgID   string

	EmptyMessagesLimit uint
}

func DefaultConfig(authToken string) ClientConfig {
	uri, err := url.Parse("http://127.0.0.1:7890")
	if err != nil {
		log.Fatalln(err)
	}

	return ClientConfig{
		HTTPClient: &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(uri),
			},
		},
		BaseURL:   apiURLv1,
		OrgID:     "",
		authToken: authToken,

		EmptyMessagesLimit: defaultEmptyMessagesLimit,
	}
}
