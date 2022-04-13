package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateAlarmRuleStatusRequest struct {
	Body *ChangeAlarmRuleStatus `json:"body,omitempty"`
}

func (o UpdateAlarmRuleStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlarmRuleStatusRequest struct{}"
	}

	return strings.Join([]string{"UpdateAlarmRuleStatusRequest", string(data)}, " ")
}
