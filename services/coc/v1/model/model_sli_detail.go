package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SliDetail Sli的详细信息
type SliDetail struct {

	// SLi的ID
	Id *string `json:"id,omitempty"`

	// 顺序
	SortId *int32 `json:"sort_id,omitempty"`

	// SLI类型 REQUEST 请求型SLI指标 INSTANCES 实例型SLI指标
	SliType SliDetailSliType `json:"sli_type"`

	// SLI名称
	Name string `json:"name"`

	// SLI描述
	Description string `json:"description"`

	// 比较符 LESS_THAN 小于 LESS_THAN_OR_EQUAL_TO 小于等于 EQUALS 等于 GREATER_THAN 大于 GREATER_THAN_OR_EQUAL_TO 大于等于
	ComparisonOperator SliDetailComparisonOperator `json:"comparison_operator"`

	// 数值
	NumericalValue float64 `json:"numerical_value"`

	// 单位 PERCENT_SIGN 百分号 MILLISECONDS 毫秒 NUMBER_OF_REQUESTS_PER_SECOND 每秒请求数量
	Unit SliDetailUnit `json:"unit"`
}

func (o SliDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SliDetail struct{}"
	}

	return strings.Join([]string{"SliDetail", string(data)}, " ")
}

type SliDetailSliType struct {
	value string
}

type SliDetailSliTypeEnum struct {
	REQUEST   SliDetailSliType
	INSTANCES SliDetailSliType
}

func GetSliDetailSliTypeEnum() SliDetailSliTypeEnum {
	return SliDetailSliTypeEnum{
		REQUEST: SliDetailSliType{
			value: "REQUEST",
		},
		INSTANCES: SliDetailSliType{
			value: "INSTANCES",
		},
	}
}

func (c SliDetailSliType) Value() string {
	return c.value
}

func (c SliDetailSliType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SliDetailSliType) UnmarshalJSON(b []byte) error {
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

type SliDetailComparisonOperator struct {
	value string
}

type SliDetailComparisonOperatorEnum struct {
	GREATER_THAN             SliDetailComparisonOperator
	GREATER_THAN_OR_EQUAL_TO SliDetailComparisonOperator
	EQUALS                   SliDetailComparisonOperator
	LESS_THAN                SliDetailComparisonOperator
	LESS_THAN_OR_EQUAL_TO    SliDetailComparisonOperator
}

func GetSliDetailComparisonOperatorEnum() SliDetailComparisonOperatorEnum {
	return SliDetailComparisonOperatorEnum{
		GREATER_THAN: SliDetailComparisonOperator{
			value: "GREATER_THAN",
		},
		GREATER_THAN_OR_EQUAL_TO: SliDetailComparisonOperator{
			value: "GREATER_THAN_OR_EQUAL_TO",
		},
		EQUALS: SliDetailComparisonOperator{
			value: "EQUALS",
		},
		LESS_THAN: SliDetailComparisonOperator{
			value: "LESS_THAN",
		},
		LESS_THAN_OR_EQUAL_TO: SliDetailComparisonOperator{
			value: "LESS_THAN_OR_EQUAL_TO",
		},
	}
}

func (c SliDetailComparisonOperator) Value() string {
	return c.value
}

func (c SliDetailComparisonOperator) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SliDetailComparisonOperator) UnmarshalJSON(b []byte) error {
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

type SliDetailUnit struct {
	value string
}

type SliDetailUnitEnum struct {
	PERCENT_SIGN                  SliDetailUnit
	MILLISECONDS                  SliDetailUnit
	NUMBER_OF_REQUESTS_PER_SECOND SliDetailUnit
}

func GetSliDetailUnitEnum() SliDetailUnitEnum {
	return SliDetailUnitEnum{
		PERCENT_SIGN: SliDetailUnit{
			value: "PERCENT_SIGN",
		},
		MILLISECONDS: SliDetailUnit{
			value: "MILLISECONDS",
		},
		NUMBER_OF_REQUESTS_PER_SECOND: SliDetailUnit{
			value: "NUMBER_OF_REQUESTS_PER_SECOND",
		},
	}
}

func (c SliDetailUnit) Value() string {
	return c.value
}

func (c SliDetailUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SliDetailUnit) UnmarshalJSON(b []byte) error {
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
