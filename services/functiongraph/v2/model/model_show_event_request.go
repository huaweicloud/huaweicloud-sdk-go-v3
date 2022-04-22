package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowEventRequest struct {

	// 事件ID。
	EventId string `json:"event_id"`

	// 函数的URN（Uniform Resource Name），唯一标识函数。
	FunctionUrn string `json:"function_urn"`
}

func (o ShowEventRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEventRequest struct{}"
	}

	return strings.Join([]string{"ShowEventRequest", string(data)}, " ")
}
