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

	question := "帮我写一首五言绝句"

	print("> ")
	println(question)
	println()

	params := deepseek.ChatCompletionNewParams{
		Messages: deepseek.F([]deepseek.ChatCompletionMessageParamUnion{
			deepseek.UserMessage(question),
		}),
		Model: deepseek.F(deepseek.ChatModelDeepSeek_Reasoner),
	}
	completion, err := client.Chat.Completions.New(ctx, params)

	if err != nil {
		panic(err)
	}

	print("output reasoning=>")
	println(completion.Choices[0].Message.ReasoningContent)
	print("output message=>")
	println(completion.Choices[0].Message.Content)

}
