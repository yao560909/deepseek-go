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

	question1 := "世界上最高的山是什么？"

	print("> ")
	println(question1)
	println()

	params := deepseek.ChatCompletionNewParams{
		Messages: deepseek.F([]deepseek.ChatCompletionMessageParamUnion{
			deepseek.UserMessage(question1),
		}),
		Model: deepseek.F(deepseek.ChatModelDeepSeek_Chat),
	}

	completion, err := client.Chat.Completions.New(ctx, params)
	if err != nil {
		panic(err)
	}
	println("Round 1")
	println(completion.Choices[0].Message.Content)
	params.Messages.Value = append(params.Messages.Value, completion.Choices[0].Message)

	question2 := "第二呢？"

	print("> ")
	println(question2)
	println()
	message := deepseek.UserMessage(question2)

	params.Messages.Value = append(params.Messages.Value, message)
	completion, err = client.Chat.Completions.New(ctx, params)
	if err != nil {
		panic(err)
	}
	println("Round 2")
	println(completion.Choices[0].Message.Content)
}
