package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateWidgetInfoProperties 额外信息，当view为bar时有效
type UpdateWidgetInfoProperties struct {

	// 聚合类型，目前只有TopN这一种类型
	Filter *UpdateWidgetInfoPropertiesFilter `json:"filter,omitempty"`

	// Top值前N个
	TopN *int32 `json:"topN,omitempty"`

	// 排序字段，asc正序，desc倒序
	Order *UpdateWidgetInfoPropertiesOrder `json:"order,omitempty"`
}

func (o UpdateWidgetInfoProperties) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWidgetInfoProperties struct{}"
	}

	return strings.Join([]string{"UpdateWidgetInfoProperties", string(data)}, " ")
}

type UpdateWidgetInfoPropertiesFilter struct {
	value string
}

type UpdateWidgetInfoPropertiesFilterEnum struct {
	TOP_N UpdateWidgetInfoPropertiesFilter
}

func GetUpdateWidgetInfoPropertiesFilterEnum() UpdateWidgetInfoPropertiesFilterEnum {
	return UpdateWidgetInfoPropertiesFilterEnum{
		TOP_N: UpdateWidgetInfoPropertiesFilter{
			value: "topN",
		},
	}
}

func (c UpdateWidgetInfoPropertiesFilter) Value() string {
	return c.value
}

func (c UpdateWidgetInfoPropertiesFilter) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateWidgetInfoPropertiesFilter) UnmarshalJSON(b []byte) error {
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

type UpdateWidgetInfoPropertiesOrder struct {
	value string
}

type UpdateWidgetInfoPropertiesOrderEnum struct {
	ASC  UpdateWidgetInfoPropertiesOrder
	DESC UpdateWidgetInfoPropertiesOrder
}

func GetUpdateWidgetInfoPropertiesOrderEnum() UpdateWidgetInfoPropertiesOrderEnum {
	return UpdateWidgetInfoPropertiesOrderEnum{
		ASC: UpdateWidgetInfoPropertiesOrder{
			value: "asc",
		},
		DESC: UpdateWidgetInfoPropertiesOrder{
			value: "desc",
		},
	}
}

func (c UpdateWidgetInfoPropertiesOrder) Value() string {
	return c.value
}

func (c UpdateWidgetInfoPropertiesOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateWidgetInfoPropertiesOrder) UnmarshalJSON(b []byte) error {
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
