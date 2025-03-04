package deepseek

import "github.com/yao560909/deepseek-go/option"

type ChatModel = string

const (
	ChatModelDeepSeek_Chat     ChatModel = "deepseek-chat"
	ChatModelDeepSeek_Reasoner ChatModel = "deepseek-reasoner"
	ChatModelDeepSeek_R1_1_5B  ChatModel = "deepseek-r1:1.5b"
)

type ChatService struct {
	Options     []option.RequestOption
	Completions *ChatCompletionService
}

func NewChatService(opts ...option.RequestOption) (r *ChatService) {
	r = &ChatService{}
	r.Options = opts
	r.Completions = NewChatCompletionService(opts...)
	return
}
