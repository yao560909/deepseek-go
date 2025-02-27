package main

import (
	"context"
	"github.com/yao560909/deepseek-go"
	"github.com/yao560909/deepseek-go/option"
)

func main() {
	apiKey := option.WithAPIKey("你的API_KEY")
	client := deepseek.NewClient(apiKey)
	ctx := context.Background()
	question := "帮我写一首四言绝句"

	print("> ")
	println(question)
	println()

	param := deepseek.ChatCompletionNewParams{
		Messages: deepseek.F([]deepseek.ChatCompletionMessageParamUnion{
			deepseek.UserMessage(question),
		}),
		Model: deepseek.F(deepseek.ChatModelDeepSeek_Chat),
	}
	completion, err := client.Chat.Completions.New(ctx, param)

	if err != nil {
		panic(err)
	}

	println(completion.Choices[0].Message.Content)

}
