package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteAlarmRuleTemplateItemResult struct {

	// 告警模板id。
	AlarmRuleTemplateName string `json:"alarm_rule_template_name"`
}

func (o DeleteAlarmRuleTemplateItemResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAlarmRuleTemplateItemResult struct{}"
	}

	return strings.Join([]string{"DeleteAlarmRuleTemplateItemResult", string(data)}, " ")
}
