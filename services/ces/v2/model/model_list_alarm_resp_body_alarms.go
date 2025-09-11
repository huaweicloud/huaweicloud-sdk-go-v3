package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ListAlarmRespBodyAlarms struct {

	// **参数解释**： 告警规则id。     **取值范围**： 以al开头，后跟22个数字或字母。
	AlarmId *string `json:"alarm_id,omitempty"`

	// **参数解释**： 告警名称。     **取值范围**： 只能包含0-9/a-z/A-Z/_/-或汉字，长度[1，128]个字符。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 告警描述。     **取值范围**： 长度为[0,256]个字符。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 查询服务的命名空间，各服务命名空间请参考“[服务命名名称](ces_03_0059.xml)”。    **取值范围**： 格式为service.item；service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_。字符串的长度为[0,32]个字符。
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**： 告警策略 **取值范围**： 最多包含100个策略。
	Policies *[]Policy `json:"policies,omitempty"`

	// **参数解释**： 资源列表，关联资源需要使用查询告警规则资源接口获取。 **取值范围**： 最多支持3000个资源。
	Resources *[]ResourcesInListResp `json:"resources,omitempty"`

	Type *AlarmTypeResp `json:"type,omitempty"`

	// **参数解释**： 是否开启告警规则。 **取值范围**： 布尔值。 - true:开启。 - false:关闭。
	Enabled *bool `json:"enabled,omitempty"`

	// **参数解释**： 是否开启告警通知。     **取值范围**： 布尔值。 - true:开启。 - false:关闭。
	NotificationEnabled *bool `json:"notification_enabled,omitempty"`

	// **参数解释**： 触发告警时，通知组/主题订阅的信息。 **取值范围**： 包含的通知信息的数量最多为10个。
	AlarmNotifications *[]NotificationResp `json:"alarm_notifications,omitempty"`

	// **参数解释**： 告警恢复时，通知组/主题订阅的信息。 **取值范围**： 包含的通知信息的数量最多为10个。
	OkNotifications *[]NotificationResp `json:"ok_notifications,omitempty"`

	// **参数解释**： 告警通知开启时间。    **取值范围**： 只能包含数字、“:”，长度为[1,64]个字符。
	NotificationBeginTime *string `json:"notification_begin_time,omitempty"`

	// **参数解释**： 告警通知关闭时间。    **取值范围**： 只能包含数字、“:”，长度为[1,64]个字符。
	NotificationEndTime *string `json:"notification_end_time,omitempty"`

	// **参数解释**： 时区，形如：\"GMT-08:00\"、\"GMT+08:00\"、\"GMT+0:00\"。    **取值范围**： 长度为[1,16]个字符。
	EffectiveTimezone *string `json:"effective_timezone,omitempty"`

	// **参数解释**： 企业项目ID。     **取值范围**： 只能包含小写字母、数字、“-”。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**： 告警规则关联告警模板ID，如果传了，告警规则关联的策略会和告警模板策略联动变化。     **取值范围**： 以at开头，只包含字母、数字，长度为[2,64]个字符。
	AlarmTemplateId *string `json:"alarm_template_id,omitempty"`

	// **参数解释**： 产品层级跨维规则需要指明的规则产品名称，一般由\"服务命名空间,服务首层维度名称\"组成，如\"SYS.ECS,instance_id\"。 **取值范围**: 长度为[0,128]个字符。
	ProductName *string `json:"product_name,omitempty"`

	// **参数解释**： 产品层级跨维规则需要指明为产品层级规则，resource_level取值为product即为云产品类型，不填或者取值为dimension则为子维度类型。 **取值范围**: 枚举值。 - product：云产品。 - dimension：子维度。
	ResourceLevel *ListAlarmRespBodyAlarmsResourceLevel `json:"resource_level,omitempty"`

	// **参数解释**： 租户标签列表 **取值范围**: 最多支持20个标签。
	Tags *[]ResourceTag `json:"tags,omitempty"`
}

func (o ListAlarmRespBodyAlarms) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmRespBodyAlarms struct{}"
	}

	return strings.Join([]string{"ListAlarmRespBodyAlarms", string(data)}, " ")
}

type ListAlarmRespBodyAlarmsResourceLevel struct {
	value string
}

type ListAlarmRespBodyAlarmsResourceLevelEnum struct {
	PRODUCT   ListAlarmRespBodyAlarmsResourceLevel
	DIMENSION ListAlarmRespBodyAlarmsResourceLevel
}

func GetListAlarmRespBodyAlarmsResourceLevelEnum() ListAlarmRespBodyAlarmsResourceLevelEnum {
	return ListAlarmRespBodyAlarmsResourceLevelEnum{
		PRODUCT: ListAlarmRespBodyAlarmsResourceLevel{
			value: "product",
		},
		DIMENSION: ListAlarmRespBodyAlarmsResourceLevel{
			value: "dimension",
		},
	}
}

func (c ListAlarmRespBodyAlarmsResourceLevel) Value() string {
	return c.value
}

func (c ListAlarmRespBodyAlarmsResourceLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAlarmRespBodyAlarmsResourceLevel) UnmarshalJSON(b []byte) error {
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
