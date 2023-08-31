package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BaseWidgetInfoProperties 额外信息，当view为bar时有效
type BaseWidgetInfoProperties struct {

	// 聚合类型，目前只有TopN这一种类型
	Filter *BaseWidgetInfoPropertiesFilter `json:"filter,omitempty"`

	// Top值前N个
	TopN *int32 `json:"topN,omitempty"`

	// 排序字段，asc正序，desc倒序
	Order *BaseWidgetInfoPropertiesOrder `json:"order,omitempty"`
}

func (o BaseWidgetInfoProperties) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaseWidgetInfoProperties struct{}"
	}

	return strings.Join([]string{"BaseWidgetInfoProperties", string(data)}, " ")
}

type BaseWidgetInfoPropertiesFilter struct {
	value string
}

type BaseWidgetInfoPropertiesFilterEnum struct {
	TOP_N BaseWidgetInfoPropertiesFilter
}

func GetBaseWidgetInfoPropertiesFilterEnum() BaseWidgetInfoPropertiesFilterEnum {
	return BaseWidgetInfoPropertiesFilterEnum{
		TOP_N: BaseWidgetInfoPropertiesFilter{
			value: "topN",
		},
	}
}

func (c BaseWidgetInfoPropertiesFilter) Value() string {
	return c.value
}

func (c BaseWidgetInfoPropertiesFilter) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BaseWidgetInfoPropertiesFilter) UnmarshalJSON(b []byte) error {
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

type BaseWidgetInfoPropertiesOrder struct {
	value string
}

type BaseWidgetInfoPropertiesOrderEnum struct {
	ASC  BaseWidgetInfoPropertiesOrder
	DESC BaseWidgetInfoPropertiesOrder
}

func GetBaseWidgetInfoPropertiesOrderEnum() BaseWidgetInfoPropertiesOrderEnum {
	return BaseWidgetInfoPropertiesOrderEnum{
		ASC: BaseWidgetInfoPropertiesOrder{
			value: "asc",
		},
		DESC: BaseWidgetInfoPropertiesOrder{
			value: "desc",
		},
	}
}

func (c BaseWidgetInfoPropertiesOrder) Value() string {
	return c.value
}

func (c BaseWidgetInfoPropertiesOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BaseWidgetInfoPropertiesOrder) UnmarshalJSON(b []byte) error {
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
