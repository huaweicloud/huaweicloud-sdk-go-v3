package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowAppInfoResponse Response Object
type ShowAppInfoResponse struct {

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
	AppType        *ShowAppInfoResponseAppType `json:"app_type,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ShowAppInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowAppInfoResponse", string(data)}, " ")
}

type ShowAppInfoResponseAppType struct {
	value string
}

type ShowAppInfoResponseAppTypeEnum struct {
	APIG      ShowAppInfoResponseAppType
	IAM       ShowAppInfoResponseAppType
	APIGW     ShowAppInfoResponseAppType
	DLM       ShowAppInfoResponseAppType
	ROMA_APIC ShowAppInfoResponseAppType
}

func GetShowAppInfoResponseAppTypeEnum() ShowAppInfoResponseAppTypeEnum {
	return ShowAppInfoResponseAppTypeEnum{
		APIG: ShowAppInfoResponseAppType{
			value: "APIG",
		},
		IAM: ShowAppInfoResponseAppType{
			value: "IAM",
		},
		APIGW: ShowAppInfoResponseAppType{
			value: "APIGW",
		},
		DLM: ShowAppInfoResponseAppType{
			value: "DLM",
		},
		ROMA_APIC: ShowAppInfoResponseAppType{
			value: "ROMA_APIC",
		},
	}
}

func (c ShowAppInfoResponseAppType) Value() string {
	return c.value
}

func (c ShowAppInfoResponseAppType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAppInfoResponseAppType) UnmarshalJSON(b []byte) error {
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
