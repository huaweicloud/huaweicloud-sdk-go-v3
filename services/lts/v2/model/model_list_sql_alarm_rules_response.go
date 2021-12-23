package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSqlAlarmRulesResponse struct {
	// SQL告警

	SqlAlarmRule   *[]SqlAlarmRuleRespList `json:"sql_alarm_rule,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListSqlAlarmRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlAlarmRulesResponse struct{}"
	}

	return strings.Join([]string{"ListSqlAlarmRulesResponse", string(data)}, " ")
}
