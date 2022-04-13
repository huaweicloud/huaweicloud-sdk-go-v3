package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListSqlAlarmRulesRequest struct {
}

func (o ListSqlAlarmRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlAlarmRulesRequest struct{}"
	}

	return strings.Join([]string{"ListSqlAlarmRulesRequest", string(data)}, " ")
}
