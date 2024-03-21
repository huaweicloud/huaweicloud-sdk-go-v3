package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFunctionVersionRequest Request Object
type CreateFunctionVersionRequest struct {

	// 函数的URN，详细解释见FunctionGraph函数模型的描述。
	FunctionUrn string `json:"function_urn"`

	// 消息体的类型（格式）
	ContentType string `json:"Content-Type"`

	Body *CreateFunctionVersionRequestBody `json:"body,omitempty"`
}

func (o CreateFunctionVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFunctionVersionRequest struct{}"
	}

	return strings.Join([]string{"CreateFunctionVersionRequest", string(data)}, " ")
}
