package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ScenarioAnnotationDto 场景标注（标识登录/注册接口及其参数与响应语义）
type ScenarioAnnotationDto struct {

	// 场景类型（LOGIN/REGISTER/CUSTOM）
	ScenarioType ScenarioAnnotationDtoScenarioType `json:"scenario_type"`

	UrlMatch *UrlMatchDto `json:"url_match"`

	RequestInspection *RequestInspectionDto `json:"request_inspection"`

	ResponseInspection *ResponseInspectionDto `json:"response_inspection,omitempty"`
}

func (o ScenarioAnnotationDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScenarioAnnotationDto struct{}"
	}

	return strings.Join([]string{"ScenarioAnnotationDto", string(data)}, " ")
}

type ScenarioAnnotationDtoScenarioType struct {
	value string
}

type ScenarioAnnotationDtoScenarioTypeEnum struct {
	LOGIN    ScenarioAnnotationDtoScenarioType
	REGISTER ScenarioAnnotationDtoScenarioType
	CUSTOM   ScenarioAnnotationDtoScenarioType
}

func GetScenarioAnnotationDtoScenarioTypeEnum() ScenarioAnnotationDtoScenarioTypeEnum {
	return ScenarioAnnotationDtoScenarioTypeEnum{
		LOGIN: ScenarioAnnotationDtoScenarioType{
			value: "LOGIN",
		},
		REGISTER: ScenarioAnnotationDtoScenarioType{
			value: "REGISTER",
		},
		CUSTOM: ScenarioAnnotationDtoScenarioType{
			value: "CUSTOM",
		},
	}
}

func (c ScenarioAnnotationDtoScenarioType) Value() string {
	return c.value
}

func (c ScenarioAnnotationDtoScenarioType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScenarioAnnotationDtoScenarioType) UnmarshalJSON(b []byte) error {
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
