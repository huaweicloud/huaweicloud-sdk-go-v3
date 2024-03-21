package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVersionAliasRequest Request Object
type CreateVersionAliasRequest struct {

	// 函数的URN，详细解释见FunctionGraph函数模型的描述。
	FunctionUrn string `json:"function_urn"`

	// 消息体的类型（格式）
	ContentType string `json:"Content-Type"`

	Body *CreateVersionAliasRequestBody `json:"body,omitempty"`
}

func (o CreateVersionAliasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVersionAliasRequest struct{}"
	}

	return strings.Join([]string{"CreateVersionAliasRequest", string(data)}, " ")
}
