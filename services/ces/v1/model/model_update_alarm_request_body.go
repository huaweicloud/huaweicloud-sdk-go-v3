package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

//
type UpdateAlarmRequestBody struct {
	// 告警名称，只能包含0-9/a-z/A-Z/_/-或汉字。

	AlarmName *string `json:"alarm_name,omitempty"`
	// 告警描述，长度0-256。

	AlarmDescription *string `json:"alarm_description,omitempty"`

	Condition *Condition `json:"condition,omitempty"`
	// 是否启用该条告警触发的动作，默认为true。注：若alarm_action_enabled为true，对应的alarm_actions、ok_actions至少有一个不能为空。若alarm_actions、ok_actions同时存在时，notificationList值保持一致。

	AlarmActionEnabled *bool `json:"alarm_action_enabled,omitempty"`
	// 告警级别，默认为2，级别为1、2、3、4。分别对应紧急、重要、次要、提示。

	AlarmLevel *int32 `json:"alarm_level,omitempty"`
	// 告警类型，支持的枚举类型：EVENT.SYS：针对系统事件的告警规则；EVENT.CUSTOM：针对自定义事件的告警规则；RESOURCE_GROUP：针对资源分组的告警规则。

	AlarmType *UpdateAlarmRequestBodyAlarmType `json:"alarm_type,omitempty"`
	// 告警触发的动作。 结构样例如下： { \"type\": \"notification\",\"notificationList\": [\"urn:smn:southchina:68438a86d98e427e907e0097b7e35d47:sd\"] } type取值： notification：通知。 autoscaling：弹性伸缩。

	AlarmActions *[]AlarmActions `json:"alarm_actions,omitempty"`
	// 数据不足触发的动作（该参数已废弃，建议无需配置）。

	InsufficientdataActions *[]AlarmActions `json:"insufficientdata_actions,omitempty"`
	// 告警恢复触发的动作

	OkActions *[]AlarmActions `json:"ok_actions,omitempty"`
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
	}
}

func (c UpdateAlarmRequestBodyAlarmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAlarmRequestBodyAlarmType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
