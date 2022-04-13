package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteSqlAlarmRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSqlAlarmRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSqlAlarmRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteSqlAlarmRuleResponse", string(data)}, " ")
}
