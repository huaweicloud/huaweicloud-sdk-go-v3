package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type PostAlarmsReqV2 struct {

	// **参数解释**： 告警名称。 **约束限制**： 不涉及。 **取值范围**： 只能包含0-9/a-z/A-Z/_/-或汉字，长度[1，128]个字符。           **默认取值**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 告警描述。     **约束限制**： 不涉及。 **取值范围**： 长度为[0,256]个字符。        **默认取值**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 查询服务的命名空间，各服务命名空间请参考“[服务命名名称](ces_03_0059.xml)”。    **约束限制**： 不涉及。 **取值范围**： 格式为service.item；service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_。字符串的长度在 0 到 32个字符之间。        **默认取值**： 不涉及。
	Namespace string `json:"namespace"`

	// **参数解释**： 资源分组ID     **约束限制**： 不涉及。  **取值范围**： 以rg开头，后跟22位由字母或数字组成的字符串。长度为[2,24]个字符。       **默认取值**： 不涉及。
	ResourceGroupId *string `json:"resource_group_id,omitempty"`

	// **参数解释**： 资源列表。 **约束限制**： 告警规则类型为全部资源、资源分组时，资源维度值传空；告警规则类型为指定资源时，资源维度值必填，可以同时指定监控多个资源。 最多可以指定1000个资源维度。
	Resources [][]Dimension `json:"resources"`

	// **参数解释**： 告警策略。 **约束限制**： 当alarm_template_id字段为空时必填，不为空时不填。最多包含50个策略。
	Policies *[]AlarmRulePolicy `json:"policies,omitempty"`

	Type *AlarmType `json:"type"`

	// **参数解释**： 触发告警时，通知组/主题订阅的信息。 **约束限制**： 不涉及。 **取值范围**： 包含的通知信息的数量最多为10个。 **默认取值**： 不涉及。
	AlarmNotifications *[]Notification `json:"alarm_notifications,omitempty"`

	// **参数解释**： 告警恢复时，通知组/主题订阅的信息。 **约束限制**： 不涉及。 **取值范围**： 包含的通知信息的数量最多为10个。 **默认取值**： 不涉及。
	OkNotifications *[]Notification `json:"ok_notifications,omitempty"`

	// **参数解释**： 告警通知开启时间。    **约束限制**： 不涉及。 **取值范围**： 只能包含数字、“:”，长度为[1,64]个字符。           **默认取值**： 不涉及。
	NotificationBeginTime *string `json:"notification_begin_time,omitempty"`

	// **参数解释**： 告警通知关闭时间。    **约束限制**： 不涉及。 **取值范围**： 只能包含数字、“:”，长度为[1,64]个字符。           **默认取值**： 不涉及。
	NotificationEndTime *string `json:"notification_end_time,omitempty"`

	// **参数解释**： 时区，形如：\"GMT-08:00\"、\"GMT+08:00\"、\"GMT+0:00\"。    **约束限制**： 不涉及。 **取值范围**： 长度为[1,16]个字符。           **默认取值**： 不涉及。
	EffectiveTimezone *string `json:"effective_timezone,omitempty"`

	// **参数解释**： 企业项目ID。     **约束限制**： 不涉及。 **取值范围**： 只能包含小写字母、数字、“-”。           **默认取值**： 不涉及。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**： 是否开启告警规则。     **约束限制**： 不涉及。 **取值范围**： 布尔值。 - true:开启。 - false:关闭。 **默认取值**： true
	Enabled bool `json:"enabled"`

	// **参数解释**： 是否开启告警通知。     **约束限制**： 不涉及。 **取值范围**： 布尔值。 - true:开启。 - false:关闭。 **默认取值**： true
	NotificationEnabled bool `json:"notification_enabled"`

	// **参数解释**： 告警规则关联告警模板ID，如果传了，告警规则关联的策略会和告警模板策略联动变化。     **约束限制**： 不涉及。 **取值范围**： 以at开头，只包含字母、数字，长度为[2,64]个字符。          **默认取值**： 不涉及。
	AlarmTemplateId *string `json:"alarm_template_id,omitempty"`

	// **参数解释**： 租户标签列表。 **约束限制**： 最多包含20个标签。
	Tags *[]ResourceTag `json:"tags,omitempty"`

	// **参数解释**： 产品层级跨维规则创建时需要指明的规则产品名称，一般由\"服务命名空间,服务首层维度名称\"组成，如\"SYS.ECS,instance_id\"。 **约束限制**： 不涉及。 **取值范围**： 长度为[0,128]个字符。          **默认取值**： 不涉及。
	ProductName *string `json:"product_name,omitempty"`

	// **参数解释**： 产品层级跨维规则创建时需要指明为产品层级规则，resource_level取值为product即为云产品类型，不填或者取值为dimension则为子维度类型。 **约束限制**： 不涉及。 **取值范围**： 枚举值。 - product：云产品。 - dimension：子维度。 **默认取值**： 子维度。
	ResourceLevel *PostAlarmsReqV2ResourceLevel `json:"resource_level,omitempty"`
}

func (o PostAlarmsReqV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostAlarmsReqV2 struct{}"
	}

	return strings.Join([]string{"PostAlarmsReqV2", string(data)}, " ")
}

type PostAlarmsReqV2ResourceLevel struct {
	value string
}

type PostAlarmsReqV2ResourceLevelEnum struct {
	PRODUCT   PostAlarmsReqV2ResourceLevel
	DIMENSION PostAlarmsReqV2ResourceLevel
}

func GetPostAlarmsReqV2ResourceLevelEnum() PostAlarmsReqV2ResourceLevelEnum {
	return PostAlarmsReqV2ResourceLevelEnum{
		PRODUCT: PostAlarmsReqV2ResourceLevel{
			value: "product",
		},
		DIMENSION: PostAlarmsReqV2ResourceLevel{
			value: "dimension",
		},
	}
}

func (c PostAlarmsReqV2ResourceLevel) Value() string {
	return c.value
}

func (c PostAlarmsReqV2ResourceLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PostAlarmsReqV2ResourceLevel) UnmarshalJSON(b []byte) error {
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
