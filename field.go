package deepseek

import "github.com/yao560909/deepseek-go/internal/param"

func F[T any](value T) param.Field[T] { return param.Field[T]{Value: value, Present: true} }
