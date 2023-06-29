package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAlarmRuleStatusRequest Request Object
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
