package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeAlarmRuleStatus 改变告警规则状态
type ChangeAlarmRuleStatus struct {

	// 告警规则ID
	AlarmRuleId string `json:"alarm_rule_id"`

	// 状态（RUNNING/STOPPING）
	Status string `json:"status"`

	// 类型
	Type string `json:"type"`
}

func (o ChangeAlarmRuleStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeAlarmRuleStatus struct{}"
	}

	return strings.Join([]string{"ChangeAlarmRuleStatus", string(data)}, " ")
}
