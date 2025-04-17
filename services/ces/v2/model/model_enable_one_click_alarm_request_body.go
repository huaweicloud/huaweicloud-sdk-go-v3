package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type EnableOneClickAlarmRequestBody struct {

	// 一键告警ID
	OneClickAlarmId string `json:"one_click_alarm_id"`

	DimensionNames *DimensionNames `json:"dimension_names"`

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
	NotificationManner *EnableOneClickAlarmRequestBodyNotificationManner `json:"notification_manner,omitempty"`

	// 关联的通知策略ID列表
	NotificationPolicyIds *[]string `json:"notification_policy_ids,omitempty"`

	// 是否以默认一键告警规则重置创建
	IsReset *bool `json:"is_reset,omitempty"`

	// 打开一键告警需要同时修改告警策略及通知(当前仅支持通知策略修改)时传递的参数
	OneClickUpdateAlarms *[]EnableOneClickAlarmRequestBodyOneClickUpdateAlarms `json:"one_click_update_alarms,omitempty"`
}

func (o EnableOneClickAlarmRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableOneClickAlarmRequestBody struct{}"
	}

	return strings.Join([]string{"EnableOneClickAlarmRequestBody", string(data)}, " ")
}

type EnableOneClickAlarmRequestBodyNotificationManner struct {
	value string
}

type EnableOneClickAlarmRequestBodyNotificationMannerEnum struct {
	NOTIFICATION_GROUP  EnableOneClickAlarmRequestBodyNotificationManner
	TOPIC_SUBSCRIPTION  EnableOneClickAlarmRequestBodyNotificationManner
	NOTIFICATION_POLICY EnableOneClickAlarmRequestBodyNotificationManner
}

func GetEnableOneClickAlarmRequestBodyNotificationMannerEnum() EnableOneClickAlarmRequestBodyNotificationMannerEnum {
	return EnableOneClickAlarmRequestBodyNotificationMannerEnum{
		NOTIFICATION_GROUP: EnableOneClickAlarmRequestBodyNotificationManner{
			value: "NOTIFICATION_GROUP",
		},
		TOPIC_SUBSCRIPTION: EnableOneClickAlarmRequestBodyNotificationManner{
			value: "TOPIC_SUBSCRIPTION",
		},
		NOTIFICATION_POLICY: EnableOneClickAlarmRequestBodyNotificationManner{
			value: "NOTIFICATION_POLICY",
		},
	}
}

func (c EnableOneClickAlarmRequestBodyNotificationManner) Value() string {
	return c.value
}

func (c EnableOneClickAlarmRequestBodyNotificationManner) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EnableOneClickAlarmRequestBodyNotificationManner) UnmarshalJSON(b []byte) error {
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
