package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateOneClickAlarmNotificationsRequest Request Object
type UpdateOneClickAlarmNotificationsRequest struct {

	// 一键告警ID
	OneClickAlarmId string `json:"one_click_alarm_id"`

	Body *PutAlarmNotificationReq `json:"body,omitempty"`
}

func (o UpdateOneClickAlarmNotificationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOneClickAlarmNotificationsRequest struct{}"
	}

	return strings.Join([]string{"UpdateOneClickAlarmNotificationsRequest", string(data)}, " ")
}
