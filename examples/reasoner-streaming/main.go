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

	question := "帮我写一首七言律诗"

	print("> ")
	println(question)
	println()

	params := deepseek.ChatCompletionNewParams{
		Messages: deepseek.F([]deepseek.ChatCompletionMessageParamUnion{
			deepseek.UserMessage(question),
		}),
		Model: deepseek.F(deepseek.ChatModelDeepSeek_Reasoner),
	}
	stream := client.Chat.Completions.NewStreaming(ctx, params)
	for stream.Next() {
		evt := stream.Current()
		if len(evt.Choices) > 0 {
			if evt.Choices[0].Delta.ReasoningContent != "" {
				print(evt.Choices[0].Delta.ReasoningContent)
			} else {
				print(evt.Choices[0].Delta.Content)
			}
		}
	}
	println()

	if err := stream.Err(); err != nil {
		panic(err)
	}

}
