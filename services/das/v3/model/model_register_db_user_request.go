package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type RegisterDbUserRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 语言
	XLanguage *RegisterDbUserRequestXLanguage `json:"X-Language,omitempty" xml:"X-Language"`

	Body *RegisterDbUserRequestBody `json:"body,omitempty" xml:"body"`
}

func (o RegisterDbUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterDbUserRequest struct{}"
	}

	return strings.Join([]string{"RegisterDbUserRequest", string(data)}, " ")
}

type RegisterDbUserRequestXLanguage struct {
	value string
}

type RegisterDbUserRequestXLanguageEnum struct {
	ZH_CN RegisterDbUserRequestXLanguage
	EN_US RegisterDbUserRequestXLanguage
}

func GetRegisterDbUserRequestXLanguageEnum() RegisterDbUserRequestXLanguageEnum {
	return RegisterDbUserRequestXLanguageEnum{
		ZH_CN: RegisterDbUserRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: RegisterDbUserRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c RegisterDbUserRequestXLanguage) Value() string {
	return c.value
}

func (c RegisterDbUserRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RegisterDbUserRequestXLanguage) UnmarshalJSON(b []byte) error {
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
