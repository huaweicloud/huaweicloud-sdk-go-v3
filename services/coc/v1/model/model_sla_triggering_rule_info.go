package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SlaTriggeringRuleInfo SLA触发规则
type SlaTriggeringRuleInfo struct {

	// 天
	Day *int32 `json:"day,omitempty"`

	// 小时
	Hour *int32 `json:"hour,omitempty"`

	// 分钟
	Minute *int32 `json:"minute,omitempty"`

	// 通知类型
	Protocol *string `json:"protocol,omitempty"`

	// 是否持续通知
	IsOngoingNotification *bool `json:"is_ongoing_notification,omitempty"`

	// 通知间隔
	NotificationInterval *int32 `json:"notification_interval,omitempty"`

	// 通知配置
	NotificationObjConfiguration *[]NotificationObjConfiguration `json:"notification_obj_configuration,omitempty"`
}

func (o SlaTriggeringRuleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlaTriggeringRuleInfo struct{}"
	}

	return strings.Join([]string{"SlaTriggeringRuleInfo", string(data)}, " ")
}
