package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/yao560909/deepseek-go"
	"github.com/yao560909/deepseek-go/option"
	"io"
	"net/http"
)

var loggingMiddleware = func(r *http.Request, mn option.MiddlewareNext) (*http.Response, error) {
	fmt.Printf("Received request: %s %s", r.Method, r.URL.Path)
	fmt.Println()
	var bodyBytes []byte
	if r.Body != nil {
		bodyBytes, _ = io.ReadAll(r.Body)
		fmt.Printf("Request Body: %s", string(bodyBytes))
		fmt.Println()
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}
	return mn(r)
}

func main() {
	apiKey := option.WithAPIKey("你的API_KEY")
	middleware := option.WithMiddleware(loggingMiddleware)
	client := deepseek.NewClient(apiKey, middleware)
	ctx := context.Background()

	question := "帮我写一首七言律诗"

	print("> ")
	println(question)
	println()

	param := deepseek.ChatCompletionNewParams{
		Messages: deepseek.F([]deepseek.ChatCompletionMessageParamUnion{
			deepseek.UserMessage(question),
		}),
		Model: deepseek.F(deepseek.ChatModelDeepSeek_Chat),
	}

	stream := client.Chat.Completions.NewStreaming(ctx, param)

	for stream.Next() {
		evt := stream.Current()
		if len(evt.Choices) > 0 {
			print(evt.Choices[0].Delta.Content)
		}
	}
	println()

	if err := stream.Err(); err != nil {
		panic(err)
	}

}
