package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type EnableOneClickAlarmRequestBodyOneClickUpdateAlarms struct {

	// **参数解释**： 告警规则id。     **约束限制**： 不涉及。 **取值范围**： 以al开头，后跟22个数字或字母。           **默认取值**： 不涉及。
	AlarmId *string `json:"alarm_id,omitempty"`

	// **参数解释**： 此字段已废弃。告警名称。      **约束限制**： 不涉及。 **取值范围**： 只能包含0-9/a-z/A-Z/_/-或汉字，长度[1，128]个字符。           **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 此字段已废弃。告警描述。     **约束限制**： 不涉及。 **取值范围**： 长度为[0,256]个字符。        **默认取值**： 不涉及。
	Description *string `json:"description,omitempty"`

	// 查询服务的命名空间，各服务命名空间请参考“[服务命名空间](ces_03_0059.xml)”
	Namespace *string `json:"namespace,omitempty"`

	// 告警策略
	Policies *[]Policy `json:"policies,omitempty"`

	// 此字段已废弃。资源列表，关联资源需要使用查询告警规则资源接口获取
	Resources *[][]Dimension `json:"resources,omitempty"`

	// **参数解释**： 此字段已废弃。告警规则类型 **约束限制**： 不涉及。 **取值范围**： 枚举值。 - ALL_INSTANCE：全部资源指标告警。 - RESOURCE_GROUP：资源分组指标告警。 - MULTI_INSTANCE：指定资源指标告警。 - EVENT.SYS：系统事件告警。 - EVENT.CUSTOM：自定义事件告警。 - DNSHealthCheck：健康检查告警。 **默认取值**： 不涉及。
	Type *EnableOneClickAlarmRequestBodyOneClickUpdateAlarmsType `json:"type,omitempty"`

	// **参数解释**： 是否开启告警规则。     **约束限制**： 不涉及。 **取值范围**： 布尔值。 - true:开启。 - false:关闭。 **默认取值**： true
	Enabled *bool `json:"enabled,omitempty"`

	// **参数解释**： 是否开启告警通知。     **约束限制**： 不涉及。 **取值范围**： 布尔值。 - true:开启。 - false:关闭。 **默认取值**： true
	NotificationEnabled *bool `json:"notification_enabled,omitempty"`

	// **参数解释**： 触发告警时，通知组/主题订阅的信息。 **约束限制**： 不涉及。 **取值范围**： 包含的通知信息的数量最多为10个。 **默认取值**： 不涉及。
	AlarmNotifications *[]Notification `json:"alarm_notifications,omitempty"`

	// **参数解释**： 告警恢复时，通知组/主题订阅的信息。 **约束限制**： 不涉及。 **取值范围**： 包含的通知信息的数量最多为10个。 **默认取值**： 不涉及。
	OkNotifications *[]Notification `json:"ok_notifications,omitempty"`

	// **参数解释**： 告警通知开启时间。    **约束限制**： 不涉及。 **取值范围**： 只能包含数字、“:”，长度为[1,64]个字符。           **默认取值**： 不涉及。
	NotificationBeginTime *string `json:"notification_begin_time,omitempty"`

	// **参数解释**： 告警通知关闭时间。    **约束限制**： 不涉及。 **取值范围**： 只能包含数字、“:”，长度为[1,64]个字符。           **默认取值**： 不涉及。
	NotificationEndTime *string `json:"notification_end_time,omitempty"`

	// **参数解释**： 此字段已废弃。通知方式。NOTIFICATION_POLICY表示通知策略，NOTIFICATION_GROUP表示通知组，TOPIC_SUBSCRIPTION表示主题订阅。一键告警原子能力是否开启告警通知以外层notification_manner参数为准，内层字段废弃。 **约束限制**： 不涉及。 **取值范围**： 枚举值：NOTIFICATION_POLICY、NOTIFICATION_GROUP、TOPIC_SUBSCRIPTION。长度为[1,64]个字符。           **默认取值**： 不涉及。
	NotificationManner *EnableOneClickAlarmRequestBodyOneClickUpdateAlarmsNotificationManner `json:"notification_manner,omitempty"`

	// **参数解释**： 此字段已废弃。关联的通知策略ID列表。一键告警原子能力是否开启告警通知以外层notification_policy_ids参数为准，内层字段废弃。 **约束限制**： 不涉及。 **取值范围**： ^([a-z]|[A-Z]|[0-9]|-){2,64}$。长度为[2,64]个字符。           **默认取值**： 不涉及。
	NotificationPolicyIds *[]string `json:"notification_policy_ids,omitempty"`
}

func (o EnableOneClickAlarmRequestBodyOneClickUpdateAlarms) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableOneClickAlarmRequestBodyOneClickUpdateAlarms struct{}"
	}

	return strings.Join([]string{"EnableOneClickAlarmRequestBodyOneClickUpdateAlarms", string(data)}, " ")
}

type EnableOneClickAlarmRequestBodyOneClickUpdateAlarmsType struct {
	value string
}

type EnableOneClickAlarmRequestBodyOneClickUpdateAlarmsTypeEnum struct {
	EVENT_SYS        EnableOneClickAlarmRequestBodyOneClickUpdateAlarmsType
	EVENT_CUSTOM     EnableOneClickAlarmRequestBodyOneClickUpdateAlarmsType
	DNS_HEALTH_CHECK EnableOneClickAlarmRequestBodyOneClickUpdateAlarmsType
	RESOURCE_GROUP   EnableOneClickAlarmRequestBodyOneClickUpdateAlarmsType
	MULTI_INSTANCE   EnableOneClickAlarmRequestBodyOneClickUpdateAlarmsType
	ALL_INSTANCE     EnableOneClickAlarmRequestBodyOneClickUpdateAlarmsType
}

func GetEnableOneClickAlarmRequestBodyOneClickUpdateAlarmsTypeEnum() EnableOneClickAlarmRequestBodyOneClickUpdateAlarmsTypeEnum {
	return EnableOneClickAlarmRequestBodyOneClickUpdateAlarmsTypeEnum{
		EVENT_SYS: EnableOneClickAlarmRequestBodyOneClickUpdateAlarmsType{
			value: "EVENT.SYS",
		},
		EVENT_CUSTOM: EnableOneClickAlarmRequestBodyOneClickUpdateAlarmsType{
			value: "EVENT.CUSTOM",
		},
		DNS_HEALTH_CHECK: EnableOneClickAlarmRequestBodyOneClickUpdateAlarmsType{
			value: "DNSHealthCheck",
		},
		RESOURCE_GROUP: EnableOneClickAlarmRequestBodyOneClickUpdateAlarmsType{
			value: "RESOURCE_GROUP",
		},
		MULTI_INSTANCE: EnableOneClickAlarmRequestBodyOneClickUpdateAlarmsType{
			value: "MULTI_INSTANCE",
		},
		ALL_INSTANCE: EnableOneClickAlarmRequestBodyOneClickUpdateAlarmsType{
			value: "ALL_INSTANCE",
		},
	}
}

func (c EnableOneClickAlarmRequestBodyOneClickUpdateAlarmsType) Value() string {
	return c.value
}

func (c EnableOneClickAlarmRequestBodyOneClickUpdateAlarmsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EnableOneClickAlarmRequestBodyOneClickUpdateAlarmsType) UnmarshalJSON(b []byte) error {
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
