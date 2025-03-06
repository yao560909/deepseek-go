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
	//"请以json格式返回数据"比较关键如果没有这个提示，仅response_format字段设置为json_object是不行的
	//不支持json_schema,若想实现，需要提示语给出具体的json结构
	question := "北京面积最大的5个区，面积从大到小排列，请以json格式返回数据"

	print("> ")
	println(question)
	println()

	params := deepseek.ChatCompletionNewParams{
		Messages: deepseek.F([]deepseek.ChatCompletionMessageParamUnion{
			deepseek.UserMessage(question),
		}),
		ResponseFormat: deepseek.F[deepseek.ChatCompletionNewParamsResponseFormatUnion](
			deepseek.ResponseFormatJSONObjectParam{
				Type: deepseek.F(deepseek.ResponseFormatJSONObjectTypeJSONObject),
			},
		),
		Model: deepseek.F(deepseek.ChatModelDeepSeek_Chat),
	}
	completion, err := client.Chat.Completions.New(ctx, params)

	if err != nil {
		panic(err)
	}

	println(completion.Choices[0].Message.Content)

}
