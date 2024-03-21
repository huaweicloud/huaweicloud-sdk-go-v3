package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTracingRequest Request Object
type UpdateTracingRequest struct {

	// 函数的URN，详细解释见FunctionGraph函数模型的描述。
	FunctionUrn string `json:"function_urn"`

	// 消息体的类型（格式）
	ContentType string `json:"Content-Type"`

	Body *UpdateTracingRequestBody `json:"body,omitempty"`
}

func (o UpdateTracingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTracingRequest struct{}"
	}

	return strings.Join([]string{"UpdateTracingRequest", string(data)}, " ")
}
