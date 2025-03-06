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

	question := "北京面积最大的5个区，面积从大到小排列，给出对应的天气情况"

	print("> ")
	println(question)
	println()

	params := deepseek.ChatCompletionNewParams{
		Messages: deepseek.F([]deepseek.ChatCompletionMessageParamUnion{
			deepseek.UserMessage(question),
		}),
		Model: deepseek.F(deepseek.ChatModelDeepSeek_Chat),
		Tools: deepseek.F([]deepseek.ChatCompletionToolParam{
			{
				Type: deepseek.F(deepseek.ChatCompletionToolTypeFunction),
				Function: deepseek.F(deepseek.FunctionDefinitionParam{
					Name:        deepseek.String("get_live_weather"),
					Description: deepseek.String("Get weather at the given location"),
					Parameters: deepseek.F(deepseek.FunctionParameters{
						"type": "object",
						"properties": map[string]interface{}{
							"location": map[string]string{
								"type": "string",
							},
						},
						"required": []string{"location"},
					}),
				}),
			},
		}),
	}

	stream := client.Chat.Completions.NewStreaming(ctx, params)

	acc := deepseek.ChatCompletionAccumulator{}

	for stream.Next() {
		chunk := stream.Current()
		acc.AddChunk(chunk)

		// When this fires, the current chunk value will not contain content data
		if content, ok := acc.JustFinishedContent(); ok {
			println("Content stream finished:", content)
			println()
		}

		if tool, ok := acc.JustFinishedToolCall(); ok {
			println("Tool call stream finished:", tool.Index, tool.Name, tool.Arguments)
			println()
		}

		// It's best to use chunks after handling JustFinished events
		if len(chunk.Choices) > 0 {
			println(chunk.Choices[0].Delta.JSON.RawJSON())
		}
	}

	if err := stream.Err(); err != nil {
		panic(err)
	}

	// After the stream is finished, acc can be used like a ChatCompletion
	_ = acc.Choices[0].Message.Content

	println("Total Tokens:", acc.Usage.TotalTokens)
	println("Finish Reason:", acc.Choices[0].FinishReason)

}
