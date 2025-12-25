package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

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

	// **参数解释**： 告警通知开启时间。如 00:00    **取值范围**： 只能包含数字、“:”，长度为[1,64]个字符。
	NotificationBeginTime *string `json:"notification_begin_time,omitempty"`

	// **参数解释**： 告警通知关闭时间。如 08:00   **取值范围**： 只能包含数字、“:”，长度为[1,64]个字符。
	NotificationEndTime *string `json:"notification_end_time,omitempty"`

	// **参数解释**： 时区，形如：\"GMT-08:00\"、\"GMT+08:00\"、\"GMT+0:00\"。    **取值范围**： 长度为[1,16]个字符。
	EffectiveTimezone *string `json:"effective_timezone,omitempty"`

	// **参数解释**： 告警的通知方式 **取值范围**： - NOTIFICATION_GROUP: 通知组 - TOPIC_SUBSCRIPTION: 主题订阅 - NOTIFICATION_POLICY：通知策略
	NotificationManner *UpdateAlarmNotificationsResponseNotificationManner `json:"notification_manner,omitempty"`

	// **参数解释**： 关联的通知策略ID列表
	NotificationPolicyIds *[]string `json:"notification_policy_ids,omitempty"`
	HttpStatusCode        int       `json:"-"`
}

func (o UpdateAlarmNotificationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlarmNotificationsResponse struct{}"
	}

	return strings.Join([]string{"UpdateAlarmNotificationsResponse", string(data)}, " ")
}

type UpdateAlarmNotificationsResponseNotificationManner struct {
	value string
}

type UpdateAlarmNotificationsResponseNotificationMannerEnum struct {
	NOTIFICATION_GROUP  UpdateAlarmNotificationsResponseNotificationManner
	TOPIC_SUBSCRIPTION  UpdateAlarmNotificationsResponseNotificationManner
	NOTIFICATION_POLICY UpdateAlarmNotificationsResponseNotificationManner
}

func GetUpdateAlarmNotificationsResponseNotificationMannerEnum() UpdateAlarmNotificationsResponseNotificationMannerEnum {
	return UpdateAlarmNotificationsResponseNotificationMannerEnum{
		NOTIFICATION_GROUP: UpdateAlarmNotificationsResponseNotificationManner{
			value: "NOTIFICATION_GROUP",
		},
		TOPIC_SUBSCRIPTION: UpdateAlarmNotificationsResponseNotificationManner{
			value: "TOPIC_SUBSCRIPTION",
		},
		NOTIFICATION_POLICY: UpdateAlarmNotificationsResponseNotificationManner{
			value: "NOTIFICATION_POLICY",
		},
	}
}

func (c UpdateAlarmNotificationsResponseNotificationManner) Value() string {
	return c.value
}

func (c UpdateAlarmNotificationsResponseNotificationManner) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAlarmNotificationsResponseNotificationManner) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
