package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// MetricMap 指标表
type MetricMap struct {

	// **参数解释**: 告警等级 - TIPS 提示 - LOW 低危 - MEDIUM 中危 - HIGH 高危 - FATAL 致命  **约束限制** 不涉及 **取值范围**: - TIPS - LOW - MEDIUM - HIGH - FATAL  **默认值** 不涉及
	Key *MetricMapKey `json:"key,omitempty"`

	// 值
	Value *int32 `json:"value,omitempty"`
}

func (o MetricMap) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricMap struct{}"
	}

	return strings.Join([]string{"MetricMap", string(data)}, " ")
}

type MetricMapKey struct {
	value string
}

type MetricMapKeyEnum struct {
	TIPS   MetricMapKey
	LOW    MetricMapKey
	MEDIUM MetricMapKey
	HIGH   MetricMapKey
	FATAL  MetricMapKey
}

func GetMetricMapKeyEnum() MetricMapKeyEnum {
	return MetricMapKeyEnum{
		TIPS: MetricMapKey{
			value: "TIPS",
		},
		LOW: MetricMapKey{
			value: "LOW",
		},
		MEDIUM: MetricMapKey{
			value: "MEDIUM",
		},
		HIGH: MetricMapKey{
			value: "HIGH",
		},
		FATAL: MetricMapKey{
			value: "FATAL",
		},
	}
}

func (c MetricMapKey) Value() string {
	return c.value
}

func (c MetricMapKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MetricMapKey) UnmarshalJSON(b []byte) error {
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
