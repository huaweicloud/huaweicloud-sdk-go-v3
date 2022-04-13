package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AsyncInvokeReservedFunctionRequest struct {
	// 函数的URN（Uniform Resource Name），唯一标识函数。

	FunctionUrn string `json:"function_urn"`
	// 执行函数请求体，为json格式。

	Body map[string]interface{} `json:"body,omitempty"`
}

func (o AsyncInvokeReservedFunctionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AsyncInvokeReservedFunctionRequest struct{}"
	}

	return strings.Join([]string{"AsyncInvokeReservedFunctionRequest", string(data)}, " ")
}
