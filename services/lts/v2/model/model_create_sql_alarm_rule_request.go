package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateSqlAlarmRuleRequest struct {
	Body *CreateSqlAlarmRuleRequestBody `json:"body,omitempty" xml:"body"`
}

func (o CreateSqlAlarmRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSqlAlarmRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateSqlAlarmRuleRequest", string(data)}, " ")
}
