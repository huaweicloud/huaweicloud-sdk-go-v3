package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SlaEscalationRuleInfo SLA上升规则
type SlaEscalationRuleInfo struct {

	// 天
	Day *int32 `json:"day,omitempty"`

	// 小时
	Hour *int32 `json:"hour,omitempty"`

	// 分钟
	Minute *int32 `json:"minute,omitempty"`

	// 顺序
	Order *int32 `json:"order,omitempty"`

	// 通知类型
	Protocol *string `json:"protocol,omitempty"`

	// 通知配置
	NotificationObjConfiguration *[]NotificationObjConfiguration `json:"notification_obj_configuration,omitempty"`
}

func (o SlaEscalationRuleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlaEscalationRuleInfo struct{}"
	}

	return strings.Join([]string{"SlaEscalationRuleInfo", string(data)}, " ")
}
