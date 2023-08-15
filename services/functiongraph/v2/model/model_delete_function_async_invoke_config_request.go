package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFunctionAsyncInvokeConfigRequest Request Object
type DeleteFunctionAsyncInvokeConfigRequest struct {

	// 函数的URN（Uniform Resource Name），唯一标识函数。
	FunctionUrn string `json:"function_urn"`
}

func (o DeleteFunctionAsyncInvokeConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFunctionAsyncInvokeConfigRequest struct{}"
	}

	return strings.Join([]string{"DeleteFunctionAsyncInvokeConfigRequest", string(data)}, " ")
}
