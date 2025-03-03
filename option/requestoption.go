package option

import (
	"bytes"
	"fmt"
	"github.com/tidwall/sjson"
	"github.com/yao560909/deepseek-go/internal/requestconfig"
	"log"
	"net/http"
	"net/url"
)

type RequestOption = func(*requestconfig.RequestConfig) error

func WithBaseURL(base string) RequestOption {
	u, err := url.Parse(base)
	if err != nil {
		log.Fatalf("failed to parse BaseURL: %s\n", err)
	}
	return func(r *requestconfig.RequestConfig) error {
		r.BaseURL = u
		return nil
	}
}

func WithEnvironmentProduction() RequestOption {
	return WithBaseURL("https://api.deepseek.com/v1/")
}
func WithAPIKey(value string) RequestOption {
	return func(r *requestconfig.RequestConfig) error {
		r.APIKey = value
		return r.Apply(WithHeader("authorization", fmt.Sprintf("Bearer %s", r.APIKey)))
	}
}

func WithJSONSet(key string, value interface{}) RequestOption {
	return func(r *requestconfig.RequestConfig) (err error) {
		if buffer, ok := r.Body.(*bytes.Buffer); ok {
			b := buffer.Bytes()
			b, err = sjson.SetBytes(b, key, value)
			if err != nil {
				return err
			}
			r.Body = bytes.NewBuffer(b)
			return nil
		}

		return fmt.Errorf("cannot use WithJSONSet on a body that is not serialized as *bytes.Buffer")
	}
}

func WithHeader(key, value string) RequestOption {
	return func(r *requestconfig.RequestConfig) error {
		r.Request.Header.Set(key, value)
		return nil
	}
}

type MiddlewareNext = func(*http.Request) (*http.Response, error)
type Middleware = func(*http.Request, MiddlewareNext) (*http.Response, error)

func WithMiddleware(middlewares ...Middleware) RequestOption {
	return func(r *requestconfig.RequestConfig) error {
		r.Middlewares = append(r.Middlewares, middlewares...)
		return nil
	}
}
