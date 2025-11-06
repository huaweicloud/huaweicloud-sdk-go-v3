package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAlarmNotificationsResponse Response Object
type UpdateAlarmNotificationsResponse struct {

	// **参数解释**： 是否开启告警通知。     **取值范围**： 布尔值。 - true:开启。 - false:关闭。
	NotificationEnabled *bool `json:"notification_enabled,omitempty"`

	// **参数解释**： 触发告警时，通知组/主题订阅的信息。
	AlarmNotifications *[]NotificationResp `json:"alarm_notifications,omitempty"`

	// **参数解释**： 告警恢复时，通知组/主题订阅的信息。
	OkNotifications *[]NotificationResp `json:"ok_notifications,omitempty"`

	// **参数解释**： 告警通知开启时间。    **取值范围**： 只能包含数字、“:”，长度为[1,64]个字符。
	NotificationBeginTime *string `json:"notification_begin_time,omitempty"`

	// **参数解释**： 告警通知关闭时间。    **取值范围**： 只能包含数字、“:”，长度为[1,64]个字符。
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
