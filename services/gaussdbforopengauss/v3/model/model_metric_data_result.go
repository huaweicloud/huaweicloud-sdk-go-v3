package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type MetricDataResult struct {

	// **参数解释**: 指标ID。 **取值范围**: 不涉及。
	Metric string `json:"metric"`

	// **参数解释** 指标类型 *取值范围* - INSTANCE：实例类型。 - NODE：节点类型。 - COMPONENT：组件类型。
	Type MetricDataResultType `json:"type"`

	// **参数解释**: 指标单位。 **取值范围**: 不涉及。
	Unit string `json:"unit"`

	// **参数解释**: 指标维度及指标值。 **取值范围**: 不涉及。
	Datapoints []DatapointResult `json:"datapoints"`

	// **参数解释**: 时间戳，例如1699495140000。 **取值范围**: 不涉及。
	Timestamps []string `json:"timestamps"`
}

func (o MetricDataResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricDataResult struct{}"
	}

	return strings.Join([]string{"MetricDataResult", string(data)}, " ")
}

type MetricDataResultType struct {
	value string
}

type MetricDataResultTypeEnum struct {
	INSTANCE  MetricDataResultType
	NODE      MetricDataResultType
	COMPONENT MetricDataResultType
}

func GetMetricDataResultTypeEnum() MetricDataResultTypeEnum {
	return MetricDataResultTypeEnum{
		INSTANCE: MetricDataResultType{
			value: "INSTANCE",
		},
		NODE: MetricDataResultType{
			value: "NODE",
		},
		COMPONENT: MetricDataResultType{
			value: "COMPONENT",
		},
	}
}

func (c MetricDataResultType) Value() string {
	return c.value
}

func (c MetricDataResultType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MetricDataResultType) UnmarshalJSON(b []byte) error {
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
