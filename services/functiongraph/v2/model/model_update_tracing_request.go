package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateTracingRequest struct {
	// 函数的URN（Uniform Resource Name），唯一标识函数。

	FunctionUrn string `json:"function_urn"`

	Body *UpdateTracingRequestBody `json:"body,omitempty"`
}

func (o UpdateTracingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTracingRequest struct{}"
	}

	return strings.Join([]string{"UpdateTracingRequest", string(data)}, " ")
}
