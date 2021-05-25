package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteFunctionAsyncInvokeConfigRequest struct {
	// 函数的URN（Uniform Resource Name），唯一标识函数。

	FunctionUrn string `json:"function_urn"`
}

func (o DeleteFunctionAsyncInvokeConfigRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteFunctionAsyncInvokeConfigRequest struct{}"
	}

	return strings.Join([]string{"DeleteFunctionAsyncInvokeConfigRequest", string(data)}, " ")
}
