package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// OrchestrationMappedParam 编排后的参数配置。
type OrchestrationMappedParam struct {

	// 编排后的请求参数名，只支持英文，数字，中划线，必须以英文开头，1-128个字符，不能与已有的参数重名，默认会透传到后端。
	MappedParamName string `json:"mapped_param_name"`

	// 编排后的参数类型，支持string和number。
	MappedParamType OrchestrationMappedParamMappedParamType `json:"mapped_param_type"`

	// 编排后的参数位置，支持query和header。
	MappedParamLocation OrchestrationMappedParamMappedParamLocation `json:"mapped_param_location"`
}

func (o OrchestrationMappedParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrchestrationMappedParam struct{}"
	}

	return strings.Join([]string{"OrchestrationMappedParam", string(data)}, " ")
}

type OrchestrationMappedParamMappedParamType struct {
	value string
}

type OrchestrationMappedParamMappedParamTypeEnum struct {
	STRING OrchestrationMappedParamMappedParamType
	NUMBER OrchestrationMappedParamMappedParamType
}

func GetOrchestrationMappedParamMappedParamTypeEnum() OrchestrationMappedParamMappedParamTypeEnum {
	return OrchestrationMappedParamMappedParamTypeEnum{
		STRING: OrchestrationMappedParamMappedParamType{
			value: "string",
		},
		NUMBER: OrchestrationMappedParamMappedParamType{
			value: "number",
		},
	}
}

func (c OrchestrationMappedParamMappedParamType) Value() string {
	return c.value
}

func (c OrchestrationMappedParamMappedParamType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OrchestrationMappedParamMappedParamType) UnmarshalJSON(b []byte) error {
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

type OrchestrationMappedParamMappedParamLocation struct {
	value string
}

type OrchestrationMappedParamMappedParamLocationEnum struct {
	QUERY  OrchestrationMappedParamMappedParamLocation
	HEADER OrchestrationMappedParamMappedParamLocation
}

func GetOrchestrationMappedParamMappedParamLocationEnum() OrchestrationMappedParamMappedParamLocationEnum {
	return OrchestrationMappedParamMappedParamLocationEnum{
		QUERY: OrchestrationMappedParamMappedParamLocation{
			value: "query",
		},
		HEADER: OrchestrationMappedParamMappedParamLocation{
			value: "header",
		},
	}
}

func (c OrchestrationMappedParamMappedParamLocation) Value() string {
	return c.value
}

func (c OrchestrationMappedParamMappedParamLocation) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OrchestrationMappedParamMappedParamLocation) UnmarshalJSON(b []byte) error {
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
