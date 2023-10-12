package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAlarmNotificationsResponse Response Object
type UpdateAlarmNotificationsResponse struct {

	// 是否开启告警通知
	NotificationEnabled *bool `json:"notification_enabled,omitempty"`

	// 告警触发的动作
	AlarmNotifications *[]Notification `json:"alarm_notifications,omitempty"`

	// 告警恢复触发的动作
	OkNotifications *[]Notification `json:"ok_notifications,omitempty"`

	// 告警通知开启时间
	NotificationBeginTime *string `json:"notification_begin_time,omitempty"`

	// 告警通知关闭时间
	NotificationEndTime *string `json:"notification_end_time,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o UpdateAlarmNotificationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlarmNotificationsResponse struct{}"
	}

	return strings.Join([]string{"UpdateAlarmNotificationsResponse", string(data)}, " ")
}
