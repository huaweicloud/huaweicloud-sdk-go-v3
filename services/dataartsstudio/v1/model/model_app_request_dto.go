package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AppRequestDto struct {

	// 应用类型
	AppType *AppRequestDtoAppType `json:"app_type,omitempty"`

	// 应用名称
	Name *string `json:"name,omitempty"`

	// 应用描述
	Description *string `json:"description,omitempty"`

	// 网关类型
	ApigType *AppRequestDtoApigType `json:"apig_type,omitempty"`

	// 网关实例编号
	ApigInstanceId *string `json:"apig_instance_id,omitempty"`
}

func (o AppRequestDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppRequestDto struct{}"
	}

	return strings.Join([]string{"AppRequestDto", string(data)}, " ")
}

type AppRequestDtoAppType struct {
	value string
}

type AppRequestDtoAppTypeEnum struct {
	APIG      AppRequestDtoAppType
	IAM       AppRequestDtoAppType
	APIGW     AppRequestDtoAppType
	DLM       AppRequestDtoAppType
	ROMA_APIC AppRequestDtoAppType
}

func GetAppRequestDtoAppTypeEnum() AppRequestDtoAppTypeEnum {
	return AppRequestDtoAppTypeEnum{
		APIG: AppRequestDtoAppType{
			value: "APIG",
		},
		IAM: AppRequestDtoAppType{
			value: "IAM",
		},
		APIGW: AppRequestDtoAppType{
			value: "APIGW",
		},
		DLM: AppRequestDtoAppType{
			value: "DLM",
		},
		ROMA_APIC: AppRequestDtoAppType{
			value: "ROMA_APIC",
		},
	}
}

func (c AppRequestDtoAppType) Value() string {
	return c.value
}

func (c AppRequestDtoAppType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AppRequestDtoAppType) UnmarshalJSON(b []byte) error {
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

type AppRequestDtoApigType struct {
	value string
}

type AppRequestDtoApigTypeEnum struct {
	APIG      AppRequestDtoApigType
	APIGW     AppRequestDtoApigType
	ROMA_APIC AppRequestDtoApigType
}

func GetAppRequestDtoApigTypeEnum() AppRequestDtoApigTypeEnum {
	return AppRequestDtoApigTypeEnum{
		APIG: AppRequestDtoApigType{
			value: "APIG",
		},
		APIGW: AppRequestDtoApigType{
			value: "APIGW",
		},
		ROMA_APIC: AppRequestDtoApigType{
			value: "ROMA_APIC",
		},
	}
}

func (c AppRequestDtoApigType) Value() string {
	return c.value
}

func (c AppRequestDtoApigType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AppRequestDtoApigType) UnmarshalJSON(b []byte) error {
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
