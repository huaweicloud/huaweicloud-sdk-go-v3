package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ApiResponsePara struct {

	// 参数名
	Name *string `json:"name,omitempty"`

	// 绑定的表字段
	Field *string `json:"field,omitempty"`

	// 参数类型
	Type *ApiResponseParaType `json:"type,omitempty"`

	// 参数描述
	Description *string `json:"description,omitempty"`

	// 参数示例值
	ExampleValue *string `json:"example_value,omitempty"`
}

func (o ApiResponsePara) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiResponsePara struct{}"
	}

	return strings.Join([]string{"ApiResponsePara", string(data)}, " ")
}

type ApiResponseParaType struct {
	value string
}

type ApiResponseParaTypeEnum struct {
	REQUEST_PARAMETER_TYPE_NUMBER ApiResponseParaType
	REQUEST_PARAMETER_TYPE_STRING ApiResponseParaType
}

func GetApiResponseParaTypeEnum() ApiResponseParaTypeEnum {
	return ApiResponseParaTypeEnum{
		REQUEST_PARAMETER_TYPE_NUMBER: ApiResponseParaType{
			value: "REQUEST_PARAMETER_TYPE_NUMBER",
		},
		REQUEST_PARAMETER_TYPE_STRING: ApiResponseParaType{
			value: "REQUEST_PARAMETER_TYPE_STRING",
		},
	}
}

func (c ApiResponseParaType) Value() string {
	return c.value
}

func (c ApiResponseParaType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiResponseParaType) UnmarshalJSON(b []byte) error {
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
