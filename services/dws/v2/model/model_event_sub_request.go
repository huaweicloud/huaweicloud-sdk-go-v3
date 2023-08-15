package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventSubRequest 创建事件订阅请求体
type EventSubRequest struct {

	// 事件订阅名称
	Name string `json:"name"`

	// 事件源类型支持cluster，backup，disaster-recovery
	SourceType *string `json:"source_type,omitempty"`

	// 事件源ID
	SourceId *string `json:"source_id,omitempty"`

	// 事件类别支持management，monitor，security，system alarm
	Category *string `json:"category,omitempty"`

	// 事件级别支持normal，warning
	Severity *string `json:"severity,omitempty"`

	// 事件标签
	Tag *string `json:"tag,omitempty"`

	// 是否开启订阅 1为开启，0为关闭
	Enable *int32 `json:"enable,omitempty"`

	// 消息通知地址
	NotificationTarget string `json:"notification_target"`

	// 消息主题名称
	NotificationTargetName string `json:"notification_target_name"`

	// 消息通知类型只支持SMN
	NotificationTargetType string `json:"notification_target_type"`

	// 时区
	TimeZone *string `json:"time_zone,omitempty"`
}

func (o EventSubRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventSubRequest struct{}"
	}

	return strings.Join([]string{"EventSubRequest", string(data)}, " ")
}
