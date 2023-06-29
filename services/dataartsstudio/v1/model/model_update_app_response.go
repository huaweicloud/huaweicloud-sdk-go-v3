package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateAppResponse Response Object
type UpdateAppResponse struct {

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
	AppType        *UpdateAppResponseAppType `json:"app_type,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o UpdateAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAppResponse struct{}"
	}

	return strings.Join([]string{"UpdateAppResponse", string(data)}, " ")
}

type UpdateAppResponseAppType struct {
	value string
}

type UpdateAppResponseAppTypeEnum struct {
	APIG      UpdateAppResponseAppType
	IAM       UpdateAppResponseAppType
	APIGW     UpdateAppResponseAppType
	DLM       UpdateAppResponseAppType
	ROMA_APIC UpdateAppResponseAppType
}

func GetUpdateAppResponseAppTypeEnum() UpdateAppResponseAppTypeEnum {
	return UpdateAppResponseAppTypeEnum{
		APIG: UpdateAppResponseAppType{
			value: "APIG",
		},
		IAM: UpdateAppResponseAppType{
			value: "IAM",
		},
		APIGW: UpdateAppResponseAppType{
			value: "APIGW",
		},
		DLM: UpdateAppResponseAppType{
			value: "DLM",
		},
		ROMA_APIC: UpdateAppResponseAppType{
			value: "ROMA_APIC",
		},
	}
}

func (c UpdateAppResponseAppType) Value() string {
	return c.value
}

func (c UpdateAppResponseAppType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAppResponseAppType) UnmarshalJSON(b []byte) error {
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
