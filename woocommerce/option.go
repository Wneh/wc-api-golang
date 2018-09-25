package woocommerce

import (
	"time"
	"net/http"
)

type Options struct {
	API             bool
	APIPrefix       string
	Version         string
	Timeout         time.Duration
	VerifySSL       bool
	QueryStringAuth string
	OauthTimestamp  time.Time
	HTTPClient *http.Client
}
