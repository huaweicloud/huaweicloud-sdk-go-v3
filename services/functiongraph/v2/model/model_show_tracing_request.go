package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowTracingRequest struct {
	// 函数的URN（Uniform Resource Name），唯一标识函数。

	FunctionUrn string `json:"function_urn"`
}

func (o ShowTracingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTracingRequest struct{}"
	}

	return strings.Join([]string{"ShowTracingRequest", string(data)}, " ")
}
