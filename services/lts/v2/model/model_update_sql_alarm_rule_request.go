package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSqlAlarmRuleRequest Request Object
type UpdateSqlAlarmRuleRequest struct {
	Body *UpdateSqlAlarmRuleRequestBody `json:"body,omitempty"`
}

func (o UpdateSqlAlarmRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSqlAlarmRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateSqlAlarmRuleRequest", string(data)}, " ")
}
