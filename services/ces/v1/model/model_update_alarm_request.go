package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAlarmRequest Request Object
type UpdateAlarmRequest struct {

	// 告警规则的ID。
	AlarmId string `json:"alarm_id"`

	Body *UpdateAlarmRequestBody `json:"body,omitempty"`
}

func (o UpdateAlarmRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlarmRequest struct{}"
	}

	return strings.Join([]string{"UpdateAlarmRequest", string(data)}, " ")
}
