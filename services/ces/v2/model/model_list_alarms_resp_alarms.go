package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ListAlarmsRespAlarms struct {

	// **参数解释**： 告警规则id。     **约束限制**： 不涉及。 **取值范围**： 以al开头，后跟22个数字或字母。           **默认取值**： 不涉及。
	AlarmId *string `json:"alarm_id,omitempty"`

	// **参数解释**： 告警名称。 **约束限制**： 不涉及。 **取值范围**： 只能包含0-9/a-z/A-Z/_/-或汉字，长度[1，128]个字符。           **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 告警描述。     **约束限制**： 不涉及。 **取值范围**： 长度为[0,256]个字符。        **默认取值**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 查询服务的命名空间，各服务命名空间请参考“[服务命名空间](ces_03_0059.xml)”。 **约束限制**： 不涉及。 **取值范围**： 格式为service.item；service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_。字符串的长度必须在 3 到 32个字符之间。 **默认取值**： 不涉及。
	Namespace *string `json:"namespace,omitempty"`

	// 告警策略
	Policies *[]OneClickAlarmPolicy `json:"policies,omitempty"`

	// 资源列表，关联资源需要使用查询告警规则资源接口获取
	Resources *[]ResourcesInListResp `json:"resources,omitempty"`

	Type *AlarmType `json:"type,omitempty"`

	// **参数解释**： 是否开启告警规则。     **约束限制**： 不涉及。 **取值范围**： 布尔值。 - true:开启。 - false:关闭。 **默认取值**： true
	Enabled *bool `json:"enabled,omitempty"`

	// **参数解释**： 是否开启告警通知。     **约束限制**： 不涉及。 **取值范围**： 布尔值。 - true:开启。 - false:关闭。 **默认取值**： true
	NotificationEnabled *bool `json:"notification_enabled,omitempty"`

	// **参数解释**： 触发告警时，通知组/主题订阅的信息。 **约束限制**： 包含的通知信息的数量最多为10个，最少为0个。
	AlarmNotifications *[]Notification `json:"alarm_notifications,omitempty"`

	// **参数解释**： 告警恢复时，通知组/主题订阅的信息。 **约束限制**： 包含的通知信息的数量最多为10个，最少为0个。
	OkNotifications *[]Notification `json:"ok_notifications,omitempty"`

	// **参数解释**： 告警通知开启时间。    **约束限制**： 不涉及。 **取值范围**： 只能包含数字、“:”，长度为[1,64]个字符。           **默认取值**： 不涉及。
	NotificationBeginTime *string `json:"notification_begin_time,omitempty"`

	// **参数解释**： 告警通知关闭时间。    **约束限制**： 不涉及。 **取值范围**： 只能包含数字、“:”，长度为[1,64]个字符。           **默认取值**： 不涉及。
	NotificationEndTime *string `json:"notification_end_time,omitempty"`

	// **参数解释**： 时区，形如：\"GMT-08:00\"、\"GMT+08:00\"、\"GMT+0:00\"。    **约束限制**： 不涉及。 **取值范围**： 长度为[1,16]个字符。           **默认取值**： 不涉及。
	EffectiveTimezone *string `json:"effective_timezone,omitempty"`

	// NOTIFICATION_GROUP(通知组)/TOPIC_SUBSCRIPTION(主题订阅)/NOTIFICATION_POLICY(通知策略)
	NotificationManner *ListAlarmsRespAlarmsNotificationManner `json:"notification_manner,omitempty"`

	// 关联的通知策略ID列表
	NotificationPolicyIds *[]string `json:"notification_policy_ids,omitempty"`
}

func (o ListAlarmsRespAlarms) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmsRespAlarms struct{}"
	}

	return strings.Join([]string{"ListAlarmsRespAlarms", string(data)}, " ")
}

type ListAlarmsRespAlarmsNotificationManner struct {
	value string
}

type ListAlarmsRespAlarmsNotificationMannerEnum struct {
	NOTIFICATION_GROUP  ListAlarmsRespAlarmsNotificationManner
	TOPIC_SUBSCRIPTION  ListAlarmsRespAlarmsNotificationManner
	NOTIFICATION_POLICY ListAlarmsRespAlarmsNotificationManner
}

func GetListAlarmsRespAlarmsNotificationMannerEnum() ListAlarmsRespAlarmsNotificationMannerEnum {
	return ListAlarmsRespAlarmsNotificationMannerEnum{
		NOTIFICATION_GROUP: ListAlarmsRespAlarmsNotificationManner{
			value: "NOTIFICATION_GROUP",
		},
		TOPIC_SUBSCRIPTION: ListAlarmsRespAlarmsNotificationManner{
			value: "TOPIC_SUBSCRIPTION",
		},
		NOTIFICATION_POLICY: ListAlarmsRespAlarmsNotificationManner{
			value: "NOTIFICATION_POLICY",
		},
	}
}

func (c ListAlarmsRespAlarmsNotificationManner) Value() string {
	return c.value
}

func (c ListAlarmsRespAlarmsNotificationManner) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAlarmsRespAlarmsNotificationManner) UnmarshalJSON(b []byte) error {
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
