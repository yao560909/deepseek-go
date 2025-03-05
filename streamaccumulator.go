package deepseek

type ChatCompletionAccumulator struct {
	// The up-to-date accumulation of model's responses
	ChatCompletion
	choiceChatCompletionStates []chatCompletionResponseState
	justFinished               chatCompletionResponseState
}

type chatCompletionResponseState struct {
	state chatCompletionResponseStateEnum
	index int
}

func (prev *chatCompletionResponseState) update(chunk ChatCompletionChunk) (justFinished chatCompletionResponseState) {
	delta := chunk.Choices[0].Delta
	new := chatCompletionResponseState{}
	switch {
	case !delta.JSON.Content.IsNull():
		new.state = contentResponseState
	case !delta.JSON.ToolCalls.IsNull():
		new.state = toolResponseState
		new.index = int(delta.ToolCalls[0].Index)
	default:
		new.state = finishedResponseState
	}

	if *prev != new {
		justFinished = *prev
	}
	*prev = new

	return justFinished
}

type chatCompletionResponseStateEnum int

const (
	emptyResponseState chatCompletionResponseStateEnum = iota
	contentResponseState
	refusalResponseState
	toolResponseState
	finishedResponseState
)

func (acc *ChatCompletionAccumulator) AddChunk(chunk ChatCompletionChunk) bool {
	acc.justFinished = chatCompletionResponseState{}
	if !acc.accumulateDelta(chunk) {
		return false
	}

	// only chunks with choices can cause finished events
	if len(chunk.Choices) == 0 {
		return true
	}

	chunkIndex := int(chunk.Choices[0].Index)
	acc.choiceChatCompletionStates = expandToFit(acc.choiceChatCompletionStates, chunkIndex)
	acc.justFinished = acc.choiceChatCompletionStates[chunkIndex].update(chunk)
	return true
}

func (acc *ChatCompletionAccumulator) JustFinishedContent() (content string, ok bool) {
	if acc.justFinished.state == contentResponseState {
		return acc.Choices[0].Message.Content, true
	}
	return "", false
}

func (acc *ChatCompletionAccumulator) JustFinishedToolCall() (toolcall FinishedChatCompletionToolCall, ok bool) {
	if acc.justFinished.state == toolResponseState {
		f := acc.Choices[0].Message.ToolCalls[acc.justFinished.index].Function
		return FinishedChatCompletionToolCall{
			Index: acc.justFinished.index,
			ChatCompletionMessageToolCallFunction: ChatCompletionMessageToolCallFunction{
				Name:      f.Name,
				Arguments: f.Arguments,
			},
		}, true
	}
	return FinishedChatCompletionToolCall{}, false
}

type FinishedChatCompletionToolCall struct {
	ChatCompletionMessageToolCallFunction
	Index int
}

func expandToFit[T any](slice []T, index int) []T {
	if index < len(slice) {
		return slice
	}
	if index < cap(slice) {
		return slice[:index+1]
	}
	newSlice := make([]T, index+1)
	copy(newSlice, slice)
	return newSlice
}

func (cc *ChatCompletion) accumulateDelta(chunk ChatCompletionChunk) bool {
	if len(cc.ID) == 0 {
		cc.ID = chunk.ID
	} else if cc.ID != chunk.ID {
		return false
	}

	for _, delta := range chunk.Choices {
		cc.Choices = expandToFit(cc.Choices, int(delta.Index))
		choice := &cc.Choices[delta.Index]

		choice.Index = delta.Index
		choice.FinishReason = ChatCompletionChoicesFinishReason(delta.FinishReason)

		if delta.Delta.Role != "" {
			choice.Message.Role = ChatCompletionMessageRole(delta.Delta.Role)
		}

		choice.Message.Content += delta.Delta.Content

		for j := range delta.Delta.ToolCalls {
			deltaTool := &delta.Delta.ToolCalls[j]

			choice.Message.ToolCalls = expandToFit(choice.Message.ToolCalls, int(deltaTool.Index))
			tool := &choice.Message.ToolCalls[deltaTool.Index]

			if deltaTool.ID != "" {
				tool.ID = deltaTool.ID
			}
			if deltaTool.Type != "" {
				tool.Type = ChatCompletionMessageToolCallType(deltaTool.Type)
			}
			tool.Function.Name += deltaTool.Function.Name
			tool.Function.Arguments += deltaTool.Function.Arguments
		}

		choice.Logprobs.Content = append(choice.Logprobs.Content, delta.Logprobs.Content...)
	}

	cc.Usage.CompletionTokens += chunk.Usage.CompletionTokens
	cc.Usage.PromptTokens += chunk.Usage.PromptTokens
	cc.Usage.TotalTokens += chunk.Usage.TotalTokens

	cc.Model = chunk.Model
	cc.Created = chunk.Created
	cc.SystemFingerprint = chunk.SystemFingerprint
	cc.Object = ChatCompletionObject(chunk.Object)

	return true
}
