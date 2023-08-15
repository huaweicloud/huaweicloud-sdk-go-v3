package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Rule struct {

	// 弹性伸缩规则的名称。  只能由字母、数字、中划线和下划线组成，并且长度为1～64个字符。  在一个节点组范围内，不允许重名。
	Name string `json:"name"`

	// 弹性伸缩规则的说明。  最大长度为1024字符。
	Description *string `json:"description,omitempty"`

	// 弹性伸缩规则的调整类型，只允许以下类型：  枚举值： - scale_out：扩容 - scale_in：缩容
	AdjustmentType RuleAdjustmentType `json:"adjustment_type"`

	// 触发弹性伸缩规则后，该集群处于冷却状态（不再执行弹性伸缩操作）的时长，单位为分钟。  取值范围[0～10080]，10080为一周的分钟数。
	CoolDownMinutes int32 `json:"cool_down_minutes"`

	// 单次调整集群节点的个数。  取值范围[1～100]
	ScalingAdjustment int32 `json:"scaling_adjustment"`

	Trigger *Trigger `json:"trigger"`
}

func (o Rule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Rule struct{}"
	}

	return strings.Join([]string{"Rule", string(data)}, " ")
}

type RuleAdjustmentType struct {
	value string
}

type RuleAdjustmentTypeEnum struct {
	SCALE_OUT RuleAdjustmentType
	SCALE_IN  RuleAdjustmentType
}

func GetRuleAdjustmentTypeEnum() RuleAdjustmentTypeEnum {
	return RuleAdjustmentTypeEnum{
		SCALE_OUT: RuleAdjustmentType{
			value: "scale_out",
		},
		SCALE_IN: RuleAdjustmentType{
			value: "scale_in",
		},
	}
}

func (c RuleAdjustmentType) Value() string {
	return c.value
}

func (c RuleAdjustmentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RuleAdjustmentType) UnmarshalJSON(b []byte) error {
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
