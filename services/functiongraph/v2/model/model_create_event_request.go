package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateEventRequest struct {

	// 函数的URN（Uniform Resource Name），唯一标识函数。
	FunctionUrn string `json:"function_urn"`

	Body *CreateEventRequestBody `json:"body,omitempty"`
}

func (o CreateEventRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEventRequest struct{}"
	}

	return strings.Join([]string{"CreateEventRequest", string(data)}, " ")
}
