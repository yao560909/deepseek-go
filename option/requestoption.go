package option

import "github.com/yao560909/deepseek-go/internal/requestconfig"

type RequestOption = func(RequestC) error
