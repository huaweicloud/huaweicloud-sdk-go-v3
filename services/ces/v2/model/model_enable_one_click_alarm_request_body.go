package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type EnableOneClickAlarmRequestBody struct {

	// **参数解释**： 一键告警ID。 **约束限制**： 不涉及。 **取值范围**： 只能为字母或者数字，字符长度为[1,64] **默认取值**： 不涉及。
	OneClickAlarmId string `json:"one_click_alarm_id"`

	DimensionNames *DimensionNames `json:"dimension_names"`

	// **参数解释**： 是否开启告警通知。说明：若notification_enabled为true，对应的alarm_notifications、ok_notifications至少有一个不能为空。    **约束限制**： 不涉及。 **取值范围**： 布尔值。 - true:开启。 - false:关闭。 **默认取值**： true
	NotificationEnabled bool `json:"notification_enabled"`

	// **参数解释**： 触发告警时，通知组/主题订阅的信息。 **约束限制**： 包含的通知对象信息的数量最多为10个，最少为0个。
	AlarmNotifications *[]Notification `json:"alarm_notifications,omitempty"`

	// **参数解释**： 告警恢复时，通知组/主题订阅的信息。 **约束限制**： 包含的通知对象信息的数量最多为10个，最少为0个。
	OkNotifications *[]Notification `json:"ok_notifications,omitempty"`

	// **参数解释**： 告警通知开启时间。如 00:00    **约束限制**： 不涉及。 **取值范围**： 只能包含数字、“:”，长度为[1,64]个字符。           **默认取值**： 不涉及。
	NotificationBeginTime *string `json:"notification_begin_time,omitempty"`

	// **参数解释**： 告警通知关闭时间。如 08:00  **约束限制**： 不涉及。 **取值范围**： 只能包含数字、“:”，长度为[1,64]个字符。           **默认取值**： 不涉及。
	NotificationEndTime *string `json:"notification_end_time,omitempty"`

	// **参数解释**： 时区，形如：\"GMT-08:00\"、\"GMT+08:00\"、\"GMT+0:00\"。    **约束限制**： 不涉及。 **取值范围**： 长度为[1,16]个字符。           **默认取值**： 不涉及。
	EffectiveTimezone *string `json:"effective_timezone,omitempty"`

	// **参数解释** 通知方式 **约束限制** 不涉及 **取值范围** 枚举值。 - NOTIFICATION_GROUP: 通知组 - TOPIC_SUBSCRIPTION: 主题订阅 - NOTIFICATION_POLICY: 通知策略 **默认取值** 不涉及
	NotificationManner *EnableOneClickAlarmRequestBodyNotificationManner `json:"notification_manner,omitempty"`

	// 关联的通知策略ID列表
	NotificationPolicyIds *[]string `json:"notification_policy_ids,omitempty"`

	// **参数解释**： 是否以默认一键告警规则重置创建 **约束限制**： 不涉及。 **取值范围**： - true: 一键告警规则重置创建 - false: 一键告警规则不重置创建 **默认取值**： true
	IsReset *bool `json:"is_reset,omitempty"`

	// **参数解释**: 开启一键告警时可选需要的开启的一键告警规则ID，默认为该服务下的所有一键告警规则ID。 **约束限制**: 数组元素个数[0,50] **取值范围**: 不涉及。 **默认取值**: 该服务下一键告警全部告警规则。
	EnabledAlarmIds *[]string `json:"enabled_alarm_ids,omitempty"`

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
