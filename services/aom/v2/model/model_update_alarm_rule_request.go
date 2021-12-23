package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateAlarmRuleRequest struct {
	Body *AlarmRuleParam `json:"body,omitempty"`
}

func (o UpdateAlarmRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlarmRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateAlarmRuleRequest", string(data)}, " ")
}
