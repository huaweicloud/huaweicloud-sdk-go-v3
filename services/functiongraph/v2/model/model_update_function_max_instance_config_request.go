package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFunctionMaxInstanceConfigRequest Request Object
type UpdateFunctionMaxInstanceConfigRequest struct {

	// 函数的URN，详细解释见FunctionGraph函数模型的描述。
	FunctionUrn string `json:"function_urn"`

	// 消息体的类型（格式）
	ContentType string `json:"Content-Type"`

	Body *UpdateFunctionMaxInstanceConfigRequestBody `json:"body,omitempty"`
}

func (o UpdateFunctionMaxInstanceConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFunctionMaxInstanceConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateFunctionMaxInstanceConfigRequest", string(data)}, " ")
}
