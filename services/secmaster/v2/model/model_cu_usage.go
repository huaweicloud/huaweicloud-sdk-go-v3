package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CuUsage 使用指标值对象
type CuUsage struct {

	// **参数解释**: 目录 - USAGE 使用  **约束限制** 不涉及 **取值范围**: - USAGE  **默认值** 不涉及
	UsageCategory *CuUsageUsageCategory `json:"usage_category,omitempty"`

	UsageMetric *UsageMetric `json:"usage_metric,omitempty"`
}

func (o CuUsage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CuUsage struct{}"
	}

	return strings.Join([]string{"CuUsage", string(data)}, " ")
}

type CuUsageUsageCategory struct {
	value string
}

type CuUsageUsageCategoryEnum struct {
	USAGE CuUsageUsageCategory
}

func GetCuUsageUsageCategoryEnum() CuUsageUsageCategoryEnum {
	return CuUsageUsageCategoryEnum{
		USAGE: CuUsageUsageCategory{
			value: "USAGE",
		},
	}
}

func (c CuUsageUsageCategory) Value() string {
	return c.value
}

func (c CuUsageUsageCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CuUsageUsageCategory) UnmarshalJSON(b []byte) error {
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
