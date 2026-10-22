package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ResponseInspectionDto 响应的检查规则
type ResponseInspectionDto struct {

	// 响应中表明成功或失败的组件（JSON/STATUS_CODE）
	ComponentType ResponseInspectionDtoComponentType `json:"component_type"`

	// JSON 路径（JSON时必选）
	JsonPath *string `json:"json_path,omitempty"`

	// 成功响应判定（可选）
	Success *[]string `json:"success,omitempty"`

	// 失败响应判定（可选）
	Fail *[]string `json:"fail,omitempty"`
}

func (o ResponseInspectionDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResponseInspectionDto struct{}"
	}

	return strings.Join([]string{"ResponseInspectionDto", string(data)}, " ")
}

type ResponseInspectionDtoComponentType struct {
	value string
}

type ResponseInspectionDtoComponentTypeEnum struct {
	JSON        ResponseInspectionDtoComponentType
	STATUS_CODE ResponseInspectionDtoComponentType
}

func GetResponseInspectionDtoComponentTypeEnum() ResponseInspectionDtoComponentTypeEnum {
	return ResponseInspectionDtoComponentTypeEnum{
		JSON: ResponseInspectionDtoComponentType{
			value: "JSON",
		},
		STATUS_CODE: ResponseInspectionDtoComponentType{
			value: "STATUS_CODE",
		},
	}
}

func (c ResponseInspectionDtoComponentType) Value() string {
	return c.value
}

func (c ResponseInspectionDtoComponentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResponseInspectionDtoComponentType) UnmarshalJSON(b []byte) error {
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
