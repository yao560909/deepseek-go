package deepseek

import "github.com/yao560909/deepseek-go/option"

type BetaService struct {
	Options     []option.RequestOption
	Completions *BetaCompletionService
}

func NewBetaService(opts ...option.RequestOption) (r *BetaService) {
	r = &BetaService{}
	r.Options = opts
	//r.Completions = NewChatCompletionService(opts...)
	return
}
