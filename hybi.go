package websocket

import (
	"net/http"
	"net/url"
)

const (
	ProtocolVersionHybi13    = 13
	ProtocolVersionHybi      = ProtocolVersionHybi13
	SupportedProtocolVersion = "13"
)

func Origin(config *Config, r *http.Request) (*url.URL, error) {
	var origin string
	switch config.Version {
	case ProtocolVersionHybi13:
		origin = r.Header.Get("Origin")
	}
	if origin == "" {
		return nil, nil
	}
	return url.ParseRequestURI(origin)
}
