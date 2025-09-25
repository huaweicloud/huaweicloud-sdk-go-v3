package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WaitEventTime struct {
	CodeWaitEventTime *CodeWaitEventTime `json:"code_wait_event_time"`

	ResourceWaitEventTime *ResourceWaitEventTime `json:"resource_wait_event_time"`
}

func (o WaitEventTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WaitEventTime struct{}"
	}

	return strings.Join([]string{"WaitEventTime", string(data)}, " ")
}
