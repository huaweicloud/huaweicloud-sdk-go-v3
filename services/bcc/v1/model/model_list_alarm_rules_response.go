package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmRulesResponse Response Object
type ListAlarmRulesResponse struct {

	// 告警规则总条数
	Count *int64 `json:"count,omitempty"`

	// 告警规则
	AlarmRules     *[]AlarmRuleEntity `json:"alarm_rules,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListAlarmRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmRulesResponse struct{}"
	}

	return strings.Join([]string{"ListAlarmRulesResponse", string(data)}, " ")
}
