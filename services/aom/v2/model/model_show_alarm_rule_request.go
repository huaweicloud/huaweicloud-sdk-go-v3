package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAlarmRuleRequest Request Object
type ShowAlarmRuleRequest struct {

	// 阈值规则ID。
	AlarmRuleId string `json:"alarm_rule_id"`
}

func (o ShowAlarmRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlarmRuleRequest struct{}"
	}

	return strings.Join([]string{"ShowAlarmRuleRequest", string(data)}, " ")
}
