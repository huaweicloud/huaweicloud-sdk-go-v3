package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EventResponseSource struct {

	// 计划事件来源类型
	Type *string `json:"type,omitempty"`

	// 主机计划事件ID
	HostScheduledEventId *string `json:"host_scheduled_event_id,omitempty"`
}

func (o EventResponseSource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventResponseSource struct{}"
	}

	return strings.Join([]string{"EventResponseSource", string(data)}, " ")
}
