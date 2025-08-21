package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Filter struct {

	// **参数解释**： 日志字段，如src_ip **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Field string `json:"field"`

	// **参数解释**： 值 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Values *[]string `json:"values,omitempty"`

	// **参数解释**： 操作符 **约束限制**： 不涉及 **取值范围**： equal 等于 not_equal 不等于 contain 包含 starts_with 以开始 **默认取值**： 不涉及
	Operator FilterOperator `json:"operator"`
}

func (o Filter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Filter struct{}"
	}

	return strings.Join([]string{"Filter", string(data)}, " ")
}

type FilterOperator struct {
	value string
}

type FilterOperatorEnum struct {
	EQUAL       FilterOperator
	NOT_EQUAL   FilterOperator
	CONTAIN     FilterOperator
	STARTS_WITH FilterOperator
}

func GetFilterOperatorEnum() FilterOperatorEnum {
	return FilterOperatorEnum{
		EQUAL: FilterOperator{
			value: "equal",
		},
		NOT_EQUAL: FilterOperator{
			value: "not_equal",
		},
		CONTAIN: FilterOperator{
			value: "contain",
		},
		STARTS_WITH: FilterOperator{
			value: "starts_with",
		},
	}
}

func (c FilterOperator) Value() string {
	return c.value
}

func (c FilterOperator) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FilterOperator) UnmarshalJSON(b []byte) error {
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
