package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmRuleTemplateResponse Response Object
type ListAlarmRuleTemplateResponse struct {

	// 告警模板列表。
	AlarmRuleTemplates *[]AlarmRuleTemplateBody `json:"alarm_rule_templates,omitempty"`

	// 告警模板总数。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAlarmRuleTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmRuleTemplateResponse struct{}"
	}

	return strings.Join([]string{"ListAlarmRuleTemplateResponse", string(data)}, " ")
}
