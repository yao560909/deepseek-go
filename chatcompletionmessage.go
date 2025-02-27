package deepseek

import "github.com/yao560909/deepseek-go/option"

type ChatCompletionMessageService struct {
	Options []option.RequestOption
}

func NewChatCompletionMessageService(opts ...option.RequestOption) (r *ChatCompletionMessageService) {
	r = &ChatCompletionMessageService{}
	r.Options = opts
	return
}
