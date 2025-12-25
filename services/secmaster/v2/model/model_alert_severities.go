package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AlertSeverities 告警严重程度
type AlertSeverities struct {

	// **参数解释**: 目录 - SEVERITY 严重性  **约束限制** 不涉及 **取值范围**: - SEVERITY  **默认值** 不涉及
	SeverityCategory *AlertSeveritiesSeverityCategory `json:"severity_category,omitempty"`

	Metric *MetricMap `json:"metric,omitempty"`
}

func (o AlertSeverities) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlertSeverities struct{}"
	}

	return strings.Join([]string{"AlertSeverities", string(data)}, " ")
}

type AlertSeveritiesSeverityCategory struct {
	value string
}

type AlertSeveritiesSeverityCategoryEnum struct {
	SEVERITY AlertSeveritiesSeverityCategory
}

func GetAlertSeveritiesSeverityCategoryEnum() AlertSeveritiesSeverityCategoryEnum {
	return AlertSeveritiesSeverityCategoryEnum{
		SEVERITY: AlertSeveritiesSeverityCategory{
			value: "SEVERITY",
		},
	}
}

func (c AlertSeveritiesSeverityCategory) Value() string {
	return c.value
}

func (c AlertSeveritiesSeverityCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlertSeveritiesSeverityCategory) UnmarshalJSON(b []byte) error {
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
