package main

import (
	"context"
	"github.com/yao560909/deepseek-go"
	"github.com/yao560909/deepseek-go/option"
)

func main() {
	apiKey := option.WithAPIKey("你的API_KEY")
	baseURL := option.WithBaseURL("https://api.deepseek.com/beta/")
	client := deepseek.NewClient(baseURL, apiKey)
	ctx := context.Background()

	question := "生成0到20之前随机数的代码"

	print("> ")
	println(question)
	println()

	prompt := "```go\n"

	params := deepseek.ChatCompletionNewParams{
		Messages: deepseek.F([]deepseek.ChatCompletionMessageParamUnion{
			deepseek.UserMessage(question),
			deepseek.AssistantMessage(prompt, true),
		}),
		Model: deepseek.F(deepseek.ChatModelDeepSeek_Chat),
		Stop: deepseek.F[deepseek.ChatCompletionNewParamsStopUnion](deepseek.ChatCompletionNewParamsStopArray{
			"```",
		}),
	}

	completion, err := client.Chat.Completions.New(ctx, params)

	if err != nil {
		panic(err)
	}

	println(completion.Choices[0].Message.Content)

}
