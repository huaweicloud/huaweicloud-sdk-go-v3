package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Event 事件信息
type Event struct {

	// 采样时间
	SampleTime *string `json:"sample_time,omitempty"`

	// 个数
	Count *int32 `json:"count,omitempty"`

	// 会话状态
	SessionStatus *string `json:"session_status,omitempty"`

	// 等待事件类型
	WaitEventType *string `json:"wait_event_type,omitempty"`

	// 等待事件名称
	WaitEventName *string `json:"wait_event_name,omitempty"`
}

func (o Event) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Event struct{}"
	}

	return strings.Join([]string{"Event", string(data)}, " ")
}
