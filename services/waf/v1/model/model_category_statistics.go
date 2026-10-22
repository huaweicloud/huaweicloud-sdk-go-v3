package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CategoryStatistics 分类别统计数据
type CategoryStatistics struct {

	// 防护类型
	CategoryType CategoryStatisticsCategoryType `json:"category_type"`

	// 累计检测数
	TotalCount int64 `json:"total_count"`

	// 检出风险数
	RiskCount int64 `json:"risk_count"`
}

func (o CategoryStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CategoryStatistics struct{}"
	}

	return strings.Join([]string{"CategoryStatistics", string(data)}, " ")
}

type CategoryStatisticsCategoryType struct {
	value string
}

type CategoryStatisticsCategoryTypeEnum struct {
	CONTENT_COMPLIANCE CategoryStatisticsCategoryType
	PROMPT_INJECTION   CategoryStatisticsCategoryType
	SENSITIVE_CONTENT  CategoryStatisticsCategoryType
}

func GetCategoryStatisticsCategoryTypeEnum() CategoryStatisticsCategoryTypeEnum {
	return CategoryStatisticsCategoryTypeEnum{
		CONTENT_COMPLIANCE: CategoryStatisticsCategoryType{
			value: "CONTENT_COMPLIANCE",
		},
		PROMPT_INJECTION: CategoryStatisticsCategoryType{
			value: "PROMPT_INJECTION",
		},
		SENSITIVE_CONTENT: CategoryStatisticsCategoryType{
			value: "SENSITIVE_CONTENT",
		},
	}
}

func (c CategoryStatisticsCategoryType) Value() string {
	return c.value
}

func (c CategoryStatisticsCategoryType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CategoryStatisticsCategoryType) UnmarshalJSON(b []byte) error {
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
