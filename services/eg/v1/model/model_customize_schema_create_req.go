package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CustomizeSchemaCreateReq struct {

	// 事件模型名称，租户下唯一，由字母、数字、点、下划线和中划线组成，必须字母或数字开头
	Name string `json:"name"`

	// 事件模型描述
	Description *string `json:"description,omitempty"`

	// 事件模型兼容性
	Compatibility CustomizeSchemaCreateReqCompatibility `json:"compatibility"`

	// schema内容格式
	Format *CustomizeSchemaCreateReqFormat `json:"format,omitempty"`

	// 事件示例
	DataSample *string `json:"data_sample,omitempty"`

	// 事件模型内容定义
	Definition string `json:"definition"`
}

func (o CustomizeSchemaCreateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomizeSchemaCreateReq struct{}"
	}

	return strings.Join([]string{"CustomizeSchemaCreateReq", string(data)}, " ")
}

type CustomizeSchemaCreateReqCompatibility struct {
	value string
}

type CustomizeSchemaCreateReqCompatibilityEnum struct {
	NONE     CustomizeSchemaCreateReqCompatibility
	FORWARD  CustomizeSchemaCreateReqCompatibility
	BACKWARD CustomizeSchemaCreateReqCompatibility
	FULL     CustomizeSchemaCreateReqCompatibility
}

func GetCustomizeSchemaCreateReqCompatibilityEnum() CustomizeSchemaCreateReqCompatibilityEnum {
	return CustomizeSchemaCreateReqCompatibilityEnum{
		NONE: CustomizeSchemaCreateReqCompatibility{
			value: "NONE",
		},
		FORWARD: CustomizeSchemaCreateReqCompatibility{
			value: "FORWARD",
		},
		BACKWARD: CustomizeSchemaCreateReqCompatibility{
			value: "BACKWARD",
		},
		FULL: CustomizeSchemaCreateReqCompatibility{
			value: "FULL",
		},
	}
}

func (c CustomizeSchemaCreateReqCompatibility) Value() string {
	return c.value
}

func (c CustomizeSchemaCreateReqCompatibility) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CustomizeSchemaCreateReqCompatibility) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type CustomizeSchemaCreateReqFormat struct {
	value string
}

type CustomizeSchemaCreateReqFormatEnum struct {
	JSON_SCHEMA_DRAFT_6 CustomizeSchemaCreateReqFormat
}

func GetCustomizeSchemaCreateReqFormatEnum() CustomizeSchemaCreateReqFormatEnum {
	return CustomizeSchemaCreateReqFormatEnum{
		JSON_SCHEMA_DRAFT_6: CustomizeSchemaCreateReqFormat{
			value: "JSON_SCHEMA_DRAFT_6",
		},
	}
}

func (c CustomizeSchemaCreateReqFormat) Value() string {
	return c.value
}

func (c CustomizeSchemaCreateReqFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CustomizeSchemaCreateReqFormat) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
