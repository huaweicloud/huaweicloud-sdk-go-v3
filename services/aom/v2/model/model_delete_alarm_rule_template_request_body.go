package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteAlarmRuleTemplateRequestBody struct {

	// 告警模板id列表。
	AlarmRuleTemplates []string `json:"alarm_rule_templates"`
}

func (o DeleteAlarmRuleTemplateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAlarmRuleTemplateRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteAlarmRuleTemplateRequestBody", string(data)}, " ")
}
