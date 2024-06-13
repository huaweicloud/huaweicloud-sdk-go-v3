package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlertLevel struct {

	// 告警渠道列表
	AlertChannel *[]string `json:"alert_channel,omitempty"`

	// 告警组列表
	AlertGroups *[]AlertGroup `json:"alertGroups,omitempty"`

	// 告警模板ID
	AlertTemplateId *string `json:"alertTemplateId,omitempty"`

	// 告警次数
	AlertTimes *int32 `json:"alertTimes,omitempty"`
}

func (o AlertLevel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlertLevel struct{}"
	}

	return strings.Join([]string{"AlertLevel", string(data)}, " ")
}
