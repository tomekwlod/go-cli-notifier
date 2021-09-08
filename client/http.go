package client

import (
	"net/http"
	"time"
)

type HTTPClient struct {
	client     *http.Client
	BackendURI string
}

func NewHTTPClient(uri string) HTTPClient {
	return HTTPClient{
		BackendURI: uri,
		client:     &http.Client{},
	}
}

func (c HTTPClient) Create(t, m string, d time.Duration) ([]byte, error) {
	res := []byte(`response for create reminder...`)
	return res, nil
	// the other methods in #5 '17
}
