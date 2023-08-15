package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddAlarmRuleResponse Response Object
type AddAlarmRuleResponse struct {

	// 阈值规则id。
	AlarmRuleId    *int64 `json:"alarm_rule_id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o AddAlarmRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddAlarmRuleResponse struct{}"
	}

	return strings.Join([]string{"AddAlarmRuleResponse", string(data)}, " ")
}
