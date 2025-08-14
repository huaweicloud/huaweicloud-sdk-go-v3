package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ResponseSchemaPropertiesDetailsDto 联邦Schema配置额外属性详细信息
type ResponseSchemaPropertiesDetailsDto struct {

	// 额外添加的属性的格式
	AttrNameFormat string `json:"attr_name_format"`

	// 是否包含额外属性
	Include ResponseSchemaPropertiesDetailsDtoInclude `json:"include"`
}

func (o ResponseSchemaPropertiesDetailsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResponseSchemaPropertiesDetailsDto struct{}"
	}

	return strings.Join([]string{"ResponseSchemaPropertiesDetailsDto", string(data)}, " ")
}

type ResponseSchemaPropertiesDetailsDtoInclude struct {
	value string
}

type ResponseSchemaPropertiesDetailsDtoIncludeEnum struct {
	YES ResponseSchemaPropertiesDetailsDtoInclude
}

func GetResponseSchemaPropertiesDetailsDtoIncludeEnum() ResponseSchemaPropertiesDetailsDtoIncludeEnum {
	return ResponseSchemaPropertiesDetailsDtoIncludeEnum{
		YES: ResponseSchemaPropertiesDetailsDtoInclude{
			value: "YES",
		},
	}
}

func (c ResponseSchemaPropertiesDetailsDtoInclude) Value() string {
	return c.value
}

func (c ResponseSchemaPropertiesDetailsDtoInclude) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResponseSchemaPropertiesDetailsDtoInclude) UnmarshalJSON(b []byte) error {
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
