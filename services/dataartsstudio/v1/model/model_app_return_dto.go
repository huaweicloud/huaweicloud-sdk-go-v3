package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AppReturnDto struct {

	// 应用编号
	Id *string `json:"id,omitempty"`

	// 应用名称
	Name *string `json:"name,omitempty"`

	// 应用描述
	Description *string `json:"description,omitempty"`

	// 应用key
	AppKey *string `json:"app_key,omitempty"`

	// 应用secret
	AppSecret *string `json:"app_secret,omitempty"`

	// 创建时间
	RegisterTime *int64 `json:"register_time,omitempty"`

	// 更新时间
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 创建者
	CreateUser *string `json:"create_user,omitempty"`

	// 更新者
	UpdateUser *string `json:"update_user,omitempty"`

	// 应用类型
	AppType *AppReturnDtoAppType `json:"app_type,omitempty"`
}

func (o AppReturnDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppReturnDto struct{}"
	}

	return strings.Join([]string{"AppReturnDto", string(data)}, " ")
}

type AppReturnDtoAppType struct {
	value string
}

type AppReturnDtoAppTypeEnum struct {
	APIG      AppReturnDtoAppType
	IAM       AppReturnDtoAppType
	APIGW     AppReturnDtoAppType
	DLM       AppReturnDtoAppType
	ROMA_APIC AppReturnDtoAppType
}

func GetAppReturnDtoAppTypeEnum() AppReturnDtoAppTypeEnum {
	return AppReturnDtoAppTypeEnum{
		APIG: AppReturnDtoAppType{
			value: "APIG",
		},
		IAM: AppReturnDtoAppType{
			value: "IAM",
		},
		APIGW: AppReturnDtoAppType{
			value: "APIGW",
		},
		DLM: AppReturnDtoAppType{
			value: "DLM",
		},
		ROMA_APIC: AppReturnDtoAppType{
			value: "ROMA_APIC",
		},
	}
}

func (c AppReturnDtoAppType) Value() string {
	return c.value
}

func (c AppReturnDtoAppType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AppReturnDtoAppType) UnmarshalJSON(b []byte) error {
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
