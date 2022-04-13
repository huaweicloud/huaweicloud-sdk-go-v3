package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateFunctionReservedInstancesRequest struct {
	// 函数的URN（Uniform Resource Name），唯一标识函数。

	FunctionUrn string `json:"function_urn"`

	Body *UpdateFunctionReservedInstancesRequestBody `json:"body,omitempty"`
}

func (o UpdateFunctionReservedInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFunctionReservedInstancesRequest struct{}"
	}

	return strings.Join([]string{"UpdateFunctionReservedInstancesRequest", string(data)}, " ")
}
