package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PutAlarmNotificationReq struct {

	// 是否开启告警通知。true:开启，false:关闭。
	NotificationEnabled bool `json:"notification_enabled"`

	// **参数解释**： 触发告警时，通知组/主题订阅的信息。 **约束限制**： 不涉及。 **取值范围**： 告警触发的动作数量最多为10个。 **默认取值**： 不涉及。
	AlarmNotifications *[]Notification `json:"alarm_notifications,omitempty"`

	// **参数解释**： 告警恢复时，通知组/主题订阅的信息。 **约束限制**： 不涉及。 **取值范围**： 告警恢复触发的动作数量最多为10个。 **默认取值**： 不涉及。
	OkNotifications *[]Notification `json:"ok_notifications,omitempty"`

	// **参数解释**： 每天告警通知的开始时间。 **约束限制**： 不涉及。 **取值范围**： 长度为[1,64]个字符。 **默认取值**： 不涉及。
	NotificationBeginTime *string `json:"notification_begin_time,omitempty"`

	// **参数解释**： 每天告警通知的结束时间。 **约束限制**： 不涉及。 **取值范围**： 长度为[1,64]个字符。 **默认取值**： 不涉及。
	NotificationEndTime *string `json:"notification_end_time,omitempty"`

	// 时区，形如：\"GMT-08:00\"、\"GMT+08:00\"、\"GMT+0:00\"
	EffectiveTimezone *string `json:"effective_timezone,omitempty"`
}

func (o PutAlarmNotificationReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutAlarmNotificationReq struct{}"
	}

	return strings.Join([]string{"PutAlarmNotificationReq", string(data)}, " ")
}
