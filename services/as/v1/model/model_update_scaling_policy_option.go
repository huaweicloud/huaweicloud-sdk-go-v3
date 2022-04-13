package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 修改伸缩策略
type UpdateScalingPolicyOption struct {
	// 策略名称(1-64字符)，可以用中文、字母、数字、下划线、中划线的组合。

	ScalingPolicyName *string `json:"scaling_policy_name,omitempty"`
	// 策略类型。告警策略：ALARM（与alarm_id对应）；定时策略：SCHEDULED（与scheduled_policy对应）；周期策略：RECURRENCE（与scheduled_policy对应）

	ScalingPolicyType *UpdateScalingPolicyOptionScalingPolicyType `json:"scaling_policy_type,omitempty"`
	// 告警ID，即告警规则的ID，当scaling_policy_type为ALARM时该项必选，此时scheduled_policy不生效。创建告警策略成功后，会自动为该告警ID对应的告警规则的alarm_actions字段增加类型为autoscaling的告警触发动作。告警ID通过查询云监控告警规则列表获取，详见《云监控API参考》的“查询告警规则列表”。

	AlarmId *string `json:"alarm_id,omitempty"`

	ScheduledPolicy *ScheduledPolicy `json:"scheduled_policy,omitempty"`

	ScalingPolicyAction *ScalingPolicyActionV1 `json:"scaling_policy_action,omitempty"`
	// 冷却时间，取值范围0-86400，默认为900，单位是秒。

	CoolDownTime *int32 `json:"cool_down_time,omitempty"`
}

func (o UpdateScalingPolicyOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScalingPolicyOption struct{}"
	}

	return strings.Join([]string{"UpdateScalingPolicyOption", string(data)}, " ")
}

type UpdateScalingPolicyOptionScalingPolicyType struct {
	value string
}

type UpdateScalingPolicyOptionScalingPolicyTypeEnum struct {
	ALARM      UpdateScalingPolicyOptionScalingPolicyType
	SCHEDULED  UpdateScalingPolicyOptionScalingPolicyType
	RECURRENCE UpdateScalingPolicyOptionScalingPolicyType
}

func GetUpdateScalingPolicyOptionScalingPolicyTypeEnum() UpdateScalingPolicyOptionScalingPolicyTypeEnum {
	return UpdateScalingPolicyOptionScalingPolicyTypeEnum{
		ALARM: UpdateScalingPolicyOptionScalingPolicyType{
			value: "ALARM",
		},
		SCHEDULED: UpdateScalingPolicyOptionScalingPolicyType{
			value: "SCHEDULED",
		},
		RECURRENCE: UpdateScalingPolicyOptionScalingPolicyType{
			value: "RECURRENCE",
		},
	}
}

func (c UpdateScalingPolicyOptionScalingPolicyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateScalingPolicyOptionScalingPolicyType) UnmarshalJSON(b []byte) error {
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
