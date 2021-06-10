package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Rules struct {
	// 弹性伸缩规则的名称。  只能由字母、数字、中划线和下划线组成，并且长度为1～64个字符。  在一个节点组范围内，不允许重名。

	Name string `json:"name"`
	// 弹性伸缩规则的说明。  最大长度为1024字符。

	Description *string `json:"description,omitempty"`
	// 弹性伸缩规则的调整类型，只允许以下类型：

	AdjustmentType RulesAdjustmentType `json:"adjustment_type"`
	// 触发弹性伸缩规则后，该集群处于冷却状态（不再执行弹性伸缩操作）的时长，单位为分钟。  取值范围[0～10080]，10080为一周的分钟数。

	CoolDownMinutes int32 `json:"cool_down_minutes"`
	// 单次调整集群节点的个数。  取值范围[1～100]

	ScalingAdjustment int32 `json:"scaling_adjustment"`

	Trigger *Trigger `json:"trigger"`
}

func (o Rules) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Rules struct{}"
	}

	return strings.Join([]string{"Rules", string(data)}, " ")
}

type RulesAdjustmentType struct {
	value string
}

type RulesAdjustmentTypeEnum struct {
	SCALE_OUT RulesAdjustmentType
	SCALE_IN  RulesAdjustmentType
}

func GetRulesAdjustmentTypeEnum() RulesAdjustmentTypeEnum {
	return RulesAdjustmentTypeEnum{
		SCALE_OUT: RulesAdjustmentType{
			value: "scale_out：扩容",
		},
		SCALE_IN: RulesAdjustmentType{
			value: "scale_in：缩容",
		},
	}
}

func (c RulesAdjustmentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *RulesAdjustmentType) UnmarshalJSON(b []byte) error {
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
