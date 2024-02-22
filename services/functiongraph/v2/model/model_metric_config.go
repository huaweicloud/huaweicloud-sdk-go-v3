package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// MetricConfig 流量配置
type MetricConfig struct {

	// 流量配置名称
	Name *string `json:"name,omitempty"`

	// 流量配置类型，当前只支持预留实例使用率一种类型
	Type *MetricConfigType `json:"type,omitempty"`

	// 流量阈值
	Threshold *int32 `json:"threshold,omitempty"`

	// 流量最小值
	Min *int32 `json:"min,omitempty"`
}

func (o MetricConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricConfig struct{}"
	}

	return strings.Join([]string{"MetricConfig", string(data)}, " ")
}

type MetricConfigType struct {
	value string
}

type MetricConfigTypeEnum struct {
	CONCURRENCY MetricConfigType
}

func GetMetricConfigTypeEnum() MetricConfigTypeEnum {
	return MetricConfigTypeEnum{
		CONCURRENCY: MetricConfigType{
			value: "Concurrency",
		},
	}
}

func (c MetricConfigType) Value() string {
	return c.value
}

func (c MetricConfigType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MetricConfigType) UnmarshalJSON(b []byte) error {
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
