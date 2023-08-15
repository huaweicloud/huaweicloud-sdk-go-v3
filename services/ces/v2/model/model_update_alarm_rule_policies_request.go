package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAlarmRulePoliciesRequest Request Object
type UpdateAlarmRulePoliciesRequest struct {

	// Alarm实例ID
	AlarmId string `json:"alarm_id"`

	Body *PoliciesReqV2 `json:"body,omitempty"`
}

func (o UpdateAlarmRulePoliciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlarmRulePoliciesRequest struct{}"
	}

	return strings.Join([]string{"UpdateAlarmRulePoliciesRequest", string(data)}, " ")
}
