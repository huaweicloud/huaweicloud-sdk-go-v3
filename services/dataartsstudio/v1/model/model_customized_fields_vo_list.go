package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CustomizedFieldsVoList struct {

	// 自定义项类型。 枚举值：   - TABLE: 表自定义项   - ATTRIBUTE: 属性自定义项   - SUBJECT: 主题自定义项   - METRIC: 业务指标自定义项
	Type CustomizedFieldsVoListType `json:"type"`

	// 自定义项列表。
	Fields *[]CustomizedFieldsVo `json:"fields,omitempty"`
}

func (o CustomizedFieldsVoList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomizedFieldsVoList struct{}"
	}

	return strings.Join([]string{"CustomizedFieldsVoList", string(data)}, " ")
}

type CustomizedFieldsVoListType struct {
	value string
}

type CustomizedFieldsVoListTypeEnum struct {
	TABLE     CustomizedFieldsVoListType
	ATTRIBUTE CustomizedFieldsVoListType
	SUBJECT   CustomizedFieldsVoListType
	METRIC    CustomizedFieldsVoListType
}

func GetCustomizedFieldsVoListTypeEnum() CustomizedFieldsVoListTypeEnum {
	return CustomizedFieldsVoListTypeEnum{
		TABLE: CustomizedFieldsVoListType{
			value: "TABLE",
		},
		ATTRIBUTE: CustomizedFieldsVoListType{
			value: "ATTRIBUTE",
		},
		SUBJECT: CustomizedFieldsVoListType{
			value: "SUBJECT",
		},
		METRIC: CustomizedFieldsVoListType{
			value: "METRIC",
		},
	}
}

func (c CustomizedFieldsVoListType) Value() string {
	return c.value
}

func (c CustomizedFieldsVoListType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CustomizedFieldsVoListType) UnmarshalJSON(b []byte) error {
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
