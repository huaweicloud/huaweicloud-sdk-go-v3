package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateVersionAliasRequest struct {
	// 函数的URN。

	FunctionUrn string `json:"function_urn"`

	Body *CreateVersionAliasRequestBody `json:"body,omitempty"`
}

func (o CreateVersionAliasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVersionAliasRequest struct{}"
	}

	return strings.Join([]string{"CreateVersionAliasRequest", string(data)}, " ")
}
