package requestconfig

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"time"
)

type RequestConfig struct {
	MaxRetries     int
	RequestTimeout time.Duration
	Context        context.Context
	Request        *http.Request
	BaseURL        *url.URL
	HTTPClient     *http.Client
	Middlewares    []middleware
	APIKey         string
	ResponseBodyInto interface{}
	ResponseInto **http.Response
	Body         io.Reader
}

type middleware = func(*http.Request, middlewareNext) (*http.Response, error)

type middlewareNext = func(*http.Request) (*http.Response, error)
