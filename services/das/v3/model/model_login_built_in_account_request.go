package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// LoginBuiltInAccountRequest Request Object
type LoginBuiltInAccountRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 请求语言类型。en-us：英文。 zh-cn：中文。
	XLanguage *LoginBuiltInAccountRequestXLanguage `json:"X-Language,omitempty"`

	Body *LoginBuiltInAccountRequestBody `json:"body,omitempty"`
}

func (o LoginBuiltInAccountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoginBuiltInAccountRequest struct{}"
	}

	return strings.Join([]string{"LoginBuiltInAccountRequest", string(data)}, " ")
}

type LoginBuiltInAccountRequestXLanguage struct {
	value string
}

type LoginBuiltInAccountRequestXLanguageEnum struct {
	ZH_CN LoginBuiltInAccountRequestXLanguage
	EN_US LoginBuiltInAccountRequestXLanguage
}

func GetLoginBuiltInAccountRequestXLanguageEnum() LoginBuiltInAccountRequestXLanguageEnum {
	return LoginBuiltInAccountRequestXLanguageEnum{
		ZH_CN: LoginBuiltInAccountRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: LoginBuiltInAccountRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c LoginBuiltInAccountRequestXLanguage) Value() string {
	return c.value
}

func (c LoginBuiltInAccountRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LoginBuiltInAccountRequestXLanguage) UnmarshalJSON(b []byte) error {
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
