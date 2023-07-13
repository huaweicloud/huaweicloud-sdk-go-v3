package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RequestPara struct {

	// 参数名
	Name *string `json:"name,omitempty"`

	// 参数的位置
	Position *RequestParaPosition `json:"position,omitempty"`

	// 参数的类型
	Type *RequestParaType `json:"type,omitempty"`

	// 参数的描述
	Description *string `json:"description,omitempty"`

	// 参数是否必填
	Necessary *bool `json:"necessary,omitempty"`

	// 实例值
	ExampleValue *string `json:"example_value,omitempty"`

	// 默认值
	DefaultValue *string `json:"default_value,omitempty"`
}

func (o RequestPara) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RequestPara struct{}"
	}

	return strings.Join([]string{"RequestPara", string(data)}, " ")
}

type RequestParaPosition struct {
	value string
}

type RequestParaPositionEnum struct {
	REQUEST_PARAMETER_POSITION_PATH   RequestParaPosition
	REQUEST_PARAMETER_POSITION_HEADER RequestParaPosition
	REQUEST_PARAMETER_POSITION_QUERY  RequestParaPosition
}

func GetRequestParaPositionEnum() RequestParaPositionEnum {
	return RequestParaPositionEnum{
		REQUEST_PARAMETER_POSITION_PATH: RequestParaPosition{
			value: "REQUEST_PARAMETER_POSITION_PATH",
		},
		REQUEST_PARAMETER_POSITION_HEADER: RequestParaPosition{
			value: "REQUEST_PARAMETER_POSITION_HEADER",
		},
		REQUEST_PARAMETER_POSITION_QUERY: RequestParaPosition{
			value: "REQUEST_PARAMETER_POSITION_QUERY",
		},
	}
}

func (c RequestParaPosition) Value() string {
	return c.value
}

func (c RequestParaPosition) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RequestParaPosition) UnmarshalJSON(b []byte) error {
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

type RequestParaType struct {
	value string
}

type RequestParaTypeEnum struct {
	REQUEST_PARAMETER_TYPE_NUMBER RequestParaType
	REQUEST_PARAMETER_TYPE_STRING RequestParaType
}

func GetRequestParaTypeEnum() RequestParaTypeEnum {
	return RequestParaTypeEnum{
		REQUEST_PARAMETER_TYPE_NUMBER: RequestParaType{
			value: "REQUEST_PARAMETER_TYPE_NUMBER",
		},
		REQUEST_PARAMETER_TYPE_STRING: RequestParaType{
			value: "REQUEST_PARAMETER_TYPE_STRING",
		},
	}
}

func (c RequestParaType) Value() string {
	return c.value
}

func (c RequestParaType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RequestParaType) UnmarshalJSON(b []byte) error {
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
