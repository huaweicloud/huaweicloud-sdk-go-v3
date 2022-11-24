package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateFunctionReservedInstancesCountRequest struct {

	// 函数的URN，详细解释见FunctionGraph函数模型的描述。
	FunctionUrn string `json:"function_urn"`

	Body *UpdateFunctionReservedInstancesCountRequestBody `json:"body,omitempty"`
}

func (o UpdateFunctionReservedInstancesCountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFunctionReservedInstancesCountRequest struct{}"
	}

	return strings.Join([]string{"UpdateFunctionReservedInstancesCountRequest", string(data)}, " ")
}
