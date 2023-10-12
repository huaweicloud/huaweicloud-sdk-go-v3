package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAlarmNotificationsRequest Request Object
type UpdateAlarmNotificationsRequest struct {

	// 告警规则ID
	AlarmId string `json:"alarm_id"`

	Body *PutAlarmNotificationReq `json:"body,omitempty"`
}

func (o UpdateAlarmNotificationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlarmNotificationsRequest struct{}"
	}

	return strings.Join([]string{"UpdateAlarmNotificationsRequest", string(data)}, " ")
}
