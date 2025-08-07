package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type EnableOneClickAlarmRequestBodyOneClickUpdateAlarms struct {

	// 告警规则id，以al开头，包含22个数字或字母
	AlarmId *string `json:"alarm_id,omitempty"`

	// 此字段已废弃。告警名称, 只能包含0-9/a-z/A-Z/_/-或汉字，长度1-128
	Name *string `json:"name,omitempty"`

	// 此字段已废弃。告警描述，长度0-256
	Description *string `json:"description,omitempty"`

	// 查询服务的命名空间，各服务命名空间请参考“[服务命名空间](ces_03_0059.xml)”
	Namespace *string `json:"namespace,omitempty"`

	// 告警策略
	Policies *[]Policy `json:"policies,omitempty"`

	// 此字段已废弃。资源列表，关联资源需要使用查询告警规则资源接口获取
	Resources *[][]Dimension `json:"resources,omitempty"`

	// 此字段已废弃。 **参数解释**： 告警规则类型 **约束限制**： 不涉及。 **取值范围**： 枚举值。ALL_INSTANCE为全部资源指标告警，RESOURCE_GROUP为资源分组指标告警，MULTI_INSTANCE为指定资源指标告警，EVENT.SYS为系统事件告警，EVENT.CUSTOM自定义事件告警，DNSHealthCheck为健康检查告警； **默认取值**： 不涉及。
	Type *string `json:"type,omitempty"`

	// 是否开启告警规则。true:开启，false:关闭。
	Enabled *bool `json:"enabled,omitempty"`

	// 此字段已废弃。是否开启告警通知。true:开启，false:关闭。
	NotificationEnabled *bool `json:"notification_enabled,omitempty"`

	// 此字段已废弃。告警触发的动作
	AlarmNotifications *[]Notification `json:"alarm_notifications,omitempty"`

	// 此字段已废弃。告警恢复触发的动作
	OkNotifications *[]Notification `json:"ok_notifications,omitempty"`

	// 此字段已废弃。告警通知开启时间
	NotificationBeginTime *string `json:"notification_begin_time,omitempty"`

	// 此字段已废弃。告警通知关闭时间
	NotificationEndTime *string `json:"notification_end_time,omitempty"`

	// 此字段已废弃。通知方式。NOTIFICATION_POLICY表示通知策略，NOTIFICATION_GROUP表示通知组，TOPIC_SUBSCRIPTION表示主题订阅。
	NotificationManner *EnableOneClickAlarmRequestBodyOneClickUpdateAlarmsNotificationManner `json:"notification_manner,omitempty"`

	// 此字段已废弃。关联的通知策略ID列表
	NotificationPolicyIds *[]string `json:"notification_policy_ids,omitempty"`
}

func (o EnableOneClickAlarmRequestBodyOneClickUpdateAlarms) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableOneClickAlarmRequestBodyOneClickUpdateAlarms struct{}"
	}

	return strings.Join([]string{"EnableOneClickAlarmRequestBodyOneClickUpdateAlarms", string(data)}, " ")
}

type EnableOneClickAlarmRequestBodyOneClickUpdateAlarmsNotificationManner struct {
	value string
}

type EnableOneClickAlarmRequestBodyOneClickUpdateAlarmsNotificationMannerEnum struct {
	NOTIFICATION_POLICY EnableOneClickAlarmRequestBodyOneClickUpdateAlarmsNotificationManner
	NOTIFICATION_GROUP  EnableOneClickAlarmRequestBodyOneClickUpdateAlarmsNotificationManner
	TOPIC_SUBSCRIPTION  EnableOneClickAlarmRequestBodyOneClickUpdateAlarmsNotificationManner
}

func GetEnableOneClickAlarmRequestBodyOneClickUpdateAlarmsNotificationMannerEnum() EnableOneClickAlarmRequestBodyOneClickUpdateAlarmsNotificationMannerEnum {
	return EnableOneClickAlarmRequestBodyOneClickUpdateAlarmsNotificationMannerEnum{
		NOTIFICATION_POLICY: EnableOneClickAlarmRequestBodyOneClickUpdateAlarmsNotificationManner{
			value: "NOTIFICATION_POLICY",
		},
		NOTIFICATION_GROUP: EnableOneClickAlarmRequestBodyOneClickUpdateAlarmsNotificationManner{
			value: "NOTIFICATION_GROUP",
		},
		TOPIC_SUBSCRIPTION: EnableOneClickAlarmRequestBodyOneClickUpdateAlarmsNotificationManner{
			value: "TOPIC_SUBSCRIPTION",
		},
	}
}

func (c EnableOneClickAlarmRequestBodyOneClickUpdateAlarmsNotificationManner) Value() string {
	return c.value
}

func (c EnableOneClickAlarmRequestBodyOneClickUpdateAlarmsNotificationManner) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EnableOneClickAlarmRequestBodyOneClickUpdateAlarmsNotificationManner) UnmarshalJSON(b []byte) error {
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
