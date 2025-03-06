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

type ResponseFormatJSONObjectParam struct {
	// The type of response format being defined: `json_object`
	Type param.Field[ResponseFormatJSONObjectType] `json:"type,required"`
}

func (r ResponseFormatJSONObjectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ResponseFormatJSONObjectParam) ImplementsChatCompletionNewParamsResponseFormatUnion() {}

// The type of response format being defined: `json_object`
type ResponseFormatJSONObjectType string

const (
	ResponseFormatJSONObjectTypeJSONObject ResponseFormatJSONObjectType = "json_object"
)

type ResponseFormatJSONSchemaParam struct {
	JSONSchema param.Field[ResponseFormatJSONSchemaJSONSchemaParam] `json:"json_schema,required"`
	// The type of response format being defined: `json_schema`
	Type param.Field[ResponseFormatJSONSchemaType] `json:"type,required"`
}

func (r ResponseFormatJSONSchemaParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ResponseFormatJSONSchemaParam) ImplementsChatCompletionNewParamsResponseFormatUnion() {}

type ResponseFormatJSONSchemaJSONSchemaParam struct {
	// The name of the response format. Must be a-z, A-Z, 0-9, or contain underscores
	// and dashes, with a maximum length of 64.
	Name param.Field[string] `json:"name,required"`
	// A description of what the response format is for, used by the model to determine
	// how to respond in the format.
	Description param.Field[string] `json:"description"`
	// The schema for the response format, described as a JSON Schema object.
	Schema param.Field[interface{}] `json:"schema"`
	// Whether to enable strict schema adherence when generating the output. If set to
	// true, the model will always follow the exact schema defined in the `schema`
	// field. Only a subset of JSON Schema is supported when `strict` is `true`. To
	// learn more, read the
	// [Structured Outputs guide](https://platform.openai.com/docs/guides/structured-outputs).
	Strict param.Field[bool] `json:"strict"`
}

func (r ResponseFormatJSONSchemaJSONSchemaParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of response format being defined: `json_schema`
type ResponseFormatJSONSchemaType string

const (
	ResponseFormatJSONSchemaTypeJSONSchema ResponseFormatJSONSchemaType = "json_schema"
)

type ResponseFormatTextParam struct {
	// The type of response format being defined: `text`
	Type param.Field[ResponseFormatTextType] `json:"type,required"`
}

func (r ResponseFormatTextParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ResponseFormatTextParam) ImplementsChatCompletionNewParamsResponseFormatUnion() {}

// The type of response format being defined: `text`
type ResponseFormatTextType string

const (
	ResponseFormatTextTypeText ResponseFormatTextType = "text"
)
