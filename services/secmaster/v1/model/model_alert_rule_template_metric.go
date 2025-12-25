package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AlertRuleTemplateMetric struct {

	// 活动数量
	Activity int64 `json:"activity"`

	// 类型
	Category AlertRuleTemplateMetricCategory `json:"category"`

	// 指标
	Metric map[string]float32 `json:"metric"`
}

func (o AlertRuleTemplateMetric) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlertRuleTemplateMetric struct{}"
	}

	return strings.Join([]string{"AlertRuleTemplateMetric", string(data)}, " ")
}

type AlertRuleTemplateMetricCategory struct {
	value string
}

type AlertRuleTemplateMetricCategoryEnum struct {
	GROUP_COUNT AlertRuleTemplateMetricCategory
}

func GetAlertRuleTemplateMetricCategoryEnum() AlertRuleTemplateMetricCategoryEnum {
	return AlertRuleTemplateMetricCategoryEnum{
		GROUP_COUNT: AlertRuleTemplateMetricCategory{
			value: "GROUP_COUNT",
		},
	}
}

func (c AlertRuleTemplateMetricCategory) Value() string {
	return c.value
}

func (c AlertRuleTemplateMetricCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlertRuleTemplateMetricCategory) UnmarshalJSON(b []byte) error {
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
