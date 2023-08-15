package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAlarmActionRequest Request Object
type UpdateAlarmActionRequest struct {

	// 告警规则的ID。
	AlarmId string `json:"alarm_id"`

	Body *ModifyAlarmActionReq `json:"body,omitempty"`
}

func (o UpdateAlarmActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlarmActionRequest struct{}"
	}

	return strings.Join([]string{"UpdateAlarmActionRequest", string(data)}, " ")
}
