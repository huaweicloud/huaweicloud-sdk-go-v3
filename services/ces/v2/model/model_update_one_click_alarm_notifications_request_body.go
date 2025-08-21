package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateOneClickAlarmNotificationsRequestBody struct {

	// **参数解释**： 是否开启告警通知。 **约束限制**： 值为true时，其他必选字段均需填，值为false时，其他所有字段不可选。 **取值范围**： - true：开启告警通知。 - false：不开启告警通知。 **默认取值**： 不涉及。
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

	// NOTIFICATION_GROUP(通知组)/TOPIC_SUBSCRIPTION(主题订阅)/NOTIFICATION_POLICY(通知策略)
	NotificationManner *UpdateOneClickAlarmNotificationsRequestBodyNotificationManner `json:"notification_manner,omitempty"`

	// 关联的通知策略ID列表
	NotificationPolicyIds *[]string `json:"notification_policy_ids,omitempty"`
}

func (o UpdateOneClickAlarmNotificationsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOneClickAlarmNotificationsRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateOneClickAlarmNotificationsRequestBody", string(data)}, " ")
}

type UpdateOneClickAlarmNotificationsRequestBodyNotificationManner struct {
	value string
}

type UpdateOneClickAlarmNotificationsRequestBodyNotificationMannerEnum struct {
	NOTIFICATION_GROUP  UpdateOneClickAlarmNotificationsRequestBodyNotificationManner
	TOPIC_SUBSCRIPTION  UpdateOneClickAlarmNotificationsRequestBodyNotificationManner
	NOTIFICATION_POLICY UpdateOneClickAlarmNotificationsRequestBodyNotificationManner
}

func GetUpdateOneClickAlarmNotificationsRequestBodyNotificationMannerEnum() UpdateOneClickAlarmNotificationsRequestBodyNotificationMannerEnum {
	return UpdateOneClickAlarmNotificationsRequestBodyNotificationMannerEnum{
		NOTIFICATION_GROUP: UpdateOneClickAlarmNotificationsRequestBodyNotificationManner{
			value: "NOTIFICATION_GROUP",
		},
		TOPIC_SUBSCRIPTION: UpdateOneClickAlarmNotificationsRequestBodyNotificationManner{
			value: "TOPIC_SUBSCRIPTION",
		},
		NOTIFICATION_POLICY: UpdateOneClickAlarmNotificationsRequestBodyNotificationManner{
			value: "NOTIFICATION_POLICY",
		},
	}
}

func (c UpdateOneClickAlarmNotificationsRequestBodyNotificationManner) Value() string {
	return c.value
}

func (c UpdateOneClickAlarmNotificationsRequestBodyNotificationManner) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateOneClickAlarmNotificationsRequestBodyNotificationManner) UnmarshalJSON(b []byte) error {
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
