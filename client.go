package deepseek

import (
	"github.com/yao560909/deepseek-go/option"
	"os"
)

type Client struct {
	Options []option.RequestOption
	Chat    *ChatService
}

func NewClient(opts ...option.RequestOption) (r *Client) {
	defaults := []option.RequestOption{option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("DEEPSEEK_API_KEY"); ok {
		defaults = append(defaults, option.WithAPIKey(o))
	}
	opts = append(defaults, opts...)
	r = &Client{
		Options: opts,
	}
	r.Chat = NewChatService(opts...)

	return
}
