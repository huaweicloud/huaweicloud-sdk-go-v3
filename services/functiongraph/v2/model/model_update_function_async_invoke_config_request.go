package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateFunctionAsyncInvokeConfigRequest struct {
	// 函数的URN（Uniform Resource Name），唯一标识函数。

	FunctionUrn string `json:"function_urn"`

	Body *UpdateFunctionAsyncInvokeConfigRequestBody `json:"body,omitempty"`
}

func (o UpdateFunctionAsyncInvokeConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFunctionAsyncInvokeConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateFunctionAsyncInvokeConfigRequest", string(data)}, " ")
}
