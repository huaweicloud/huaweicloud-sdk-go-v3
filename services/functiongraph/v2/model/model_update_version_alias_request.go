package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateVersionAliasRequest struct {
	// 函数的URN，详细解释见FunctionGraph函数模型的描述。

	FunctionUrn string `json:"function_urn"`
	// 要更新的别名名称。

	AliasName string `json:"alias_name"`

	Body *UpdateVersionAliasRequestBody `json:"body,omitempty"`
}

func (o UpdateVersionAliasRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateVersionAliasRequest struct{}"
	}

	return strings.Join([]string{"UpdateVersionAliasRequest", string(data)}, " ")
}
