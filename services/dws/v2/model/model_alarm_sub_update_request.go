package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlarmSubUpdateRequest struct {

	// 告警订阅名称
	Name string `json:"name"`

	// 是否开启订阅
	Enable *int32 `json:"enable,omitempty"`

	// 告警级别
	AlarmLevel *string `json:"alarm_level,omitempty"`

	// 消息主题地址
	NotificationTarget string `json:"notification_target"`

	// 消息主题名称
	NotificationTargetName string `json:"notification_target_name"`

	// 消息主题类型，支持SMN
	NotificationTargetType string `json:"notification_target_type"`
}

func (o AlarmSubUpdateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmSubUpdateRequest struct{}"
	}

	return strings.Join([]string{"AlarmSubUpdateRequest", string(data)}, " ")
}
