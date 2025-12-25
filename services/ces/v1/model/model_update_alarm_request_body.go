package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateAlarmRequestBody
type UpdateAlarmRequestBody struct {

	// **参数解释**： 告警规则名称 **约束限制**： 不涉及 **取值范围**： 只能包含0-9/a-z/A-Z/_/-或汉字，长度[1, 128]个字符 **默认取值**： 不涉及
	AlarmName *string `json:"alarm_name,omitempty"`

	// **参数解释**： 告警描述。 **约束限制**： 不涉及。 **取值范围**： 长度[0,256]个字符。 **默认取值**： 不涉及。
	AlarmDescription *string `json:"alarm_description,omitempty"`

	Condition *Condition `json:"condition,omitempty"`

	// **参数解释**： 是否开启告警通知 **约束限制**： 若alarm_action_enabled为true，对应的alarm_actions、ok_actions至少有一个不能为空。若alarm_actions、ok_actions同时存在时，notificationList值保持一致。 **取值范围**： 只能为true、false - true: 开启告警通知 - false：关闭告警通知 **默认取值**: false。
	AlarmActionEnabled *bool `json:"alarm_action_enabled,omitempty"`

	// **参数解释**： 告警级别 **约束限制**： 不涉及 **取值范围**： 级别为1、2、3、4。 - 1：紧急 - 2：重要 - 3：次要 - 4：提示 **默认取值**: 2
	AlarmLevel *int32 `json:"alarm_level,omitempty"`

	// **参数解释**： 告警类型。 **约束限制**： 不涉及 **取值范围**： - EVENT.SYS：针对系统事件的告警规则。 - EVENT.CUSTOM：针对自定义事件的告警规则。 - RESOURCE_GROUP：针对资源分组的告警规则。 - MULTI_INSTANCE： 针对指定资源的告警规则。 **默认取值**： 不涉及。
	AlarmType *UpdateAlarmRequestBodyAlarmType `json:"alarm_type,omitempty"`

	// **参数解释**： 告警触发时，通知组/主题订阅的信息。结构样例如下： { \"type\": \"notification\",\"notificationList\":[\"urn:smn:southchina:68438a86d98e427e907e0097b7e35d47:sd\"] } **约束限制**： 最多包含10个告警触发时的通知对象信息。
	AlarmActions *[][]Notification `json:"alarm_actions,omitempty"`

	// **参数解释**： 数据不足触发告警时，通知组/主题订阅的信息。（该参数已废弃，建议无需配置） **约束限制**： 最多包含10个告警动作。
	InsufficientdataActions *[][]Notification `json:"insufficientdata_actions,omitempty"`

	// **参数解释**： 告警恢复时，通知组/主题订阅的信息。结构样例如下： { \"type\": \"notification\",\"notificationList\":[\"urn:smn:southchina:68438a86d98e427e907e0097b7e35d47:sd\"] }  **约束限制**： 最多包含10个告警触发时的通知对象信息。
	OkActions *[][]Notification `json:"ok_actions,omitempty"`
}

func (o UpdateAlarmRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlarmRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAlarmRequestBody", string(data)}, " ")
}

type UpdateAlarmRequestBodyAlarmType struct {
	value string
}

type UpdateAlarmRequestBodyAlarmTypeEnum struct {
	EVENT_SYS      UpdateAlarmRequestBodyAlarmType
	EVENT_CUSTOM   UpdateAlarmRequestBodyAlarmType
	RESOURCE_GROUP UpdateAlarmRequestBodyAlarmType
	MULTI_INSTANCE UpdateAlarmRequestBodyAlarmType
}

func GetUpdateAlarmRequestBodyAlarmTypeEnum() UpdateAlarmRequestBodyAlarmTypeEnum {
	return UpdateAlarmRequestBodyAlarmTypeEnum{
		EVENT_SYS: UpdateAlarmRequestBodyAlarmType{
			value: "EVENT.SYS",
		},
		EVENT_CUSTOM: UpdateAlarmRequestBodyAlarmType{
			value: "EVENT.CUSTOM",
		},
		RESOURCE_GROUP: UpdateAlarmRequestBodyAlarmType{
			value: "RESOURCE_GROUP",
		},
		MULTI_INSTANCE: UpdateAlarmRequestBodyAlarmType{
			value: "MULTI_INSTANCE",
		},
	}
}

func (c UpdateAlarmRequestBodyAlarmType) Value() string {
	return c.value
}

func (c UpdateAlarmRequestBodyAlarmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAlarmRequestBodyAlarmType) UnmarshalJSON(b []byte) error {
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
