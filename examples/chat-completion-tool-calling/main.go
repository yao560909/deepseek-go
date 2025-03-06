package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/yao560909/deepseek-go"
	"github.com/yao560909/deepseek-go/option"
	"math/rand"
	"time"
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
					Name:        deepseek.String("get_weather"),
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
	completion, err := client.Chat.Completions.New(ctx, params)
	if err != nil {
		panic(err)
	}

	toolCalls := completion.Choices[0].Message.ToolCalls

	// Abort early if there are no tool calls
	if len(toolCalls) == 0 {
		fmt.Printf("No function call")
		return
	}

	// If there is a function call, continue the conversation
	params.Messages.Value = append(params.Messages.Value, completion.Choices[0].Message)
	for _, toolCall := range toolCalls {
		if toolCall.Function.Name == "get_weather" {
			// Extract the location from the function call arguments
			var args map[string]interface{}
			if err := json.Unmarshal([]byte(toolCall.Function.Arguments), &args); err != nil {
				panic(err)
			}
			location := args["location"].(string)

			// Simulate getting weather data
			weatherData := getWeather(location)

			// Print the weather data
			fmt.Printf("Weather in %s: %s\n", location, weatherData)

			params.Messages.Value = append(params.Messages.Value, deepseek.ToolMessage(toolCall.ID, weatherData))
		}
	}

	completion, err = client.Chat.Completions.New(ctx, params)
	if err != nil {
		panic(err)
	}

	println(completion.Choices[0].Message.Content)

}

// Mock function to simulate weather data retrieval
func getWeather(location string) string {
	// In a real implementation, this function would call a weather API
	number := randomNumber()
	return fmt.Sprintf("%s, %d°C", location, number)
}

func randomNumber() int {
	rand.New(rand.NewSource(time.Now().UnixNano())) // 设置随机数种子
	return rand.Intn(21)                            // 生成 [0, 21) 范围内的随机整数
}
