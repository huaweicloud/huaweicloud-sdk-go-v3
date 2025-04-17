package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateOneClickAlarmNotificationsRequestBody struct {

	// 是否开启告警通知
	NotificationEnabled bool `json:"notification_enabled"`

	// 告警触发的动作
	AlarmNotifications *[]Notification `json:"alarm_notifications,omitempty"`

	// 告警恢复触发的动作
	OkNotifications *[]Notification `json:"ok_notifications,omitempty"`

	// 告警通知开启时间
	NotificationBeginTime *string `json:"notification_begin_time,omitempty"`

	// 告警通知关闭时间
	NotificationEndTime *string `json:"notification_end_time,omitempty"`

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
