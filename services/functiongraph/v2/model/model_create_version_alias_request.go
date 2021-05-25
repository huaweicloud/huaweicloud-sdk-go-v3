package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateVersionAliasRequest struct {
	// 函数的URN。

	FunctionUrn string `json:"function_urn"`

	Body *CreateVersionAliasRequestBody `json:"body,omitempty"`
}

func (o CreateVersionAliasRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateVersionAliasRequest struct{}"
	}

	return strings.Join([]string{"CreateVersionAliasRequest", string(data)}, " ")
}
