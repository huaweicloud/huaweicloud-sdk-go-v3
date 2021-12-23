package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteEventRequest struct {
	// 事件ID。

	EventId string `json:"event_id"`
	// 函数的URN（Uniform Resource Name），唯一标识函数。

	FunctionUrn string `json:"function_urn"`
}

func (o DeleteEventRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEventRequest struct{}"
	}

	return strings.Join([]string{"DeleteEventRequest", string(data)}, " ")
}
