package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFunctionAsyncInvokeConfigRequest Request Object
type ShowFunctionAsyncInvokeConfigRequest struct {

	// 函数的URN（Uniform Resource Name），唯一标识函数。
	FunctionUrn string `json:"function_urn"`
}

func (o ShowFunctionAsyncInvokeConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFunctionAsyncInvokeConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowFunctionAsyncInvokeConfigRequest", string(data)}, " ")
}
