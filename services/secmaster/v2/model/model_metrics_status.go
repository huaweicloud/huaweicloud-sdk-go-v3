package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// MetricsStatus 指标状态
type MetricsStatus struct {

	// **参数解释**: 目录指标 - STATUS 状态  **约束限制** 不涉及 **取值范围**: - STATUS  **默认值** 不涉及
	MetricsCategory *MetricsStatusMetricsCategory `json:"metrics_category,omitempty"`

	StatusMetric *StatusMetric `json:"status_metric,omitempty"`
}

func (o MetricsStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricsStatus struct{}"
	}

	return strings.Join([]string{"MetricsStatus", string(data)}, " ")
}

type MetricsStatusMetricsCategory struct {
	value string
}

type MetricsStatusMetricsCategoryEnum struct {
	STATUS MetricsStatusMetricsCategory
}

func GetMetricsStatusMetricsCategoryEnum() MetricsStatusMetricsCategoryEnum {
	return MetricsStatusMetricsCategoryEnum{
		STATUS: MetricsStatusMetricsCategory{
			value: "STATUS",
		},
	}
}

func (c MetricsStatusMetricsCategory) Value() string {
	return c.value
}

func (c MetricsStatusMetricsCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MetricsStatusMetricsCategory) UnmarshalJSON(b []byte) error {
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
