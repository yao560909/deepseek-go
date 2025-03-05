package shared

import (
	"github.com/yao560909/deepseek-go/internal/apijson"
	"github.com/yao560909/deepseek-go/internal/param"
)

type FunctionDefinitionParam struct {
	Name        param.Field[string]             `json:"name,required"`
	Description param.Field[string]             `json:"description"`
	Parameters  param.Field[FunctionParameters] `json:"parameters"`
}

func (r FunctionDefinitionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FunctionParameters map[string]interface{}
