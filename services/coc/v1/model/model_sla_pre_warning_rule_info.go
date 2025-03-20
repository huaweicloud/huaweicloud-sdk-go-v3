package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SlaPreWarningRuleInfo SLA预警规则
type SlaPreWarningRuleInfo struct {

	// 天
	Day *int32 `json:"day,omitempty"`

	// 小时
	Hour *int32 `json:"hour,omitempty"`

	// 分钟
	Minute *int32 `json:"minute,omitempty"`

	// 通知类型
	Protocol *string `json:"protocol,omitempty"`

	// 通知对象配置
	NotificationObjConfiguration *[]NotificationObjConfiguration `json:"notification_obj_configuration,omitempty"`
}

func (o SlaPreWarningRuleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlaPreWarningRuleInfo struct{}"
	}

	return strings.Join([]string{"SlaPreWarningRuleInfo", string(data)}, " ")
}
