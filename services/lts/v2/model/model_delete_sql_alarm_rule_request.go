package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteSqlAlarmRuleRequest struct {

	// Sql告警规则id
	SqlAlarmRuleId string `json:"sql_alarm_rule_id" xml:"sql_alarm_rule_id"`
}

func (o DeleteSqlAlarmRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSqlAlarmRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteSqlAlarmRuleRequest", string(data)}, " ")
}
