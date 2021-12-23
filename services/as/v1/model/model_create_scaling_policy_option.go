package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 创建伸缩策略
type CreateScalingPolicyOption struct {
	// 策略名称(1-64字符)，可以用中文、字母、数字、下划线、中划线的组合。

	ScalingPolicyName string `json:"scaling_policy_name"`
	// 伸缩组ID，通过查询弹性伸缩组列表获取，详见查询弹性伸缩组列表。

	ScalingGroupId string `json:"scaling_group_id"`
	// 策略类型。告警策略：ALARM（与alarm_id对应）；定时策略：SCHEDULED（与scheduled_policy对应）；周期策略：RECURRENCE（与scheduled_policy对应）

	ScalingPolicyType CreateScalingPolicyOptionScalingPolicyType `json:"scaling_policy_type"`
	// 告警ID，即告警规则的ID，当scaling_policy_type为ALARM时该项必选，此时scheduled_policy不生效。创建告警策略成功后，会自动为该告警ID对应的告警规则的alarm_actions字段增加类型为autoscaling的告警触发动作。告警ID通过查询云监控告警规则列表获取，详见《云监控API参考》的“查询告警规则列表”。

	AlarmId *string `json:"alarm_id,omitempty"`

	ScheduledPolicy *ScheduledPolicy `json:"scheduled_policy,omitempty"`

	ScalingPolicyAction *ScalingPolicyActionV1 `json:"scaling_policy_action,omitempty"`
	// 冷却时间，取值范围0-86400，默认为900，单位是秒。

	CoolDownTime *int32 `json:"cool_down_time,omitempty"`
}

func (o CreateScalingPolicyOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScalingPolicyOption struct{}"
	}

	return strings.Join([]string{"CreateScalingPolicyOption", string(data)}, " ")
}

type CreateScalingPolicyOptionScalingPolicyType struct {
	value string
}

type CreateScalingPolicyOptionScalingPolicyTypeEnum struct {
	ALARM      CreateScalingPolicyOptionScalingPolicyType
	SCHEDULED  CreateScalingPolicyOptionScalingPolicyType
	RECURRENCE CreateScalingPolicyOptionScalingPolicyType
}

func GetCreateScalingPolicyOptionScalingPolicyTypeEnum() CreateScalingPolicyOptionScalingPolicyTypeEnum {
	return CreateScalingPolicyOptionScalingPolicyTypeEnum{
		ALARM: CreateScalingPolicyOptionScalingPolicyType{
			value: "ALARM",
		},
		SCHEDULED: CreateScalingPolicyOptionScalingPolicyType{
			value: "SCHEDULED",
		},
		RECURRENCE: CreateScalingPolicyOptionScalingPolicyType{
			value: "RECURRENCE",
		},
	}
}

func (c CreateScalingPolicyOptionScalingPolicyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateScalingPolicyOptionScalingPolicyType) UnmarshalJSON(b []byte) error {
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
