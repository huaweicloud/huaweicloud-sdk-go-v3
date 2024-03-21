package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFunctionReservedInstancesCountRequest Request Object
type UpdateFunctionReservedInstancesCountRequest struct {

	// 函数的URN，详细解释见FunctionGraph函数模型的描述。
	FunctionUrn string `json:"function_urn"`

	// 消息体的类型（格式）
	ContentType string `json:"Content-Type"`

	Body *UpdateFunctionReservedInstancesCountRequestBody `json:"body,omitempty"`
}

func (o UpdateFunctionReservedInstancesCountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFunctionReservedInstancesCountRequest struct{}"
	}

	return strings.Join([]string{"UpdateFunctionReservedInstancesCountRequest", string(data)}, " ")
}
