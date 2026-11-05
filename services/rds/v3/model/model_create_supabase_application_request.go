package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateSupabaseApplicationRequest Request Object
type CreateSupabaseApplicationRequest struct {

	// 语言
	XLanguage *CreateSupabaseApplicationRequestXLanguage `json:"X-Language,omitempty"`

	// 保证客户端请求幂等性的标识。32位UUID格式，72小时内唯一。
	XClientToken *string `json:"X-Client-Token,omitempty"`

	Body *CreateSupabaseApplicationRequestBody `json:"body,omitempty"`
}

func (o CreateSupabaseApplicationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSupabaseApplicationRequest struct{}"
	}

	return strings.Join([]string{"CreateSupabaseApplicationRequest", string(data)}, " ")
}

type CreateSupabaseApplicationRequestXLanguage struct {
	value string
}

type CreateSupabaseApplicationRequestXLanguageEnum struct {
	ZH_CN CreateSupabaseApplicationRequestXLanguage
	EN_US CreateSupabaseApplicationRequestXLanguage
}

func GetCreateSupabaseApplicationRequestXLanguageEnum() CreateSupabaseApplicationRequestXLanguageEnum {
	return CreateSupabaseApplicationRequestXLanguageEnum{
		ZH_CN: CreateSupabaseApplicationRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: CreateSupabaseApplicationRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c CreateSupabaseApplicationRequestXLanguage) Value() string {
	return c.value
}

func (c CreateSupabaseApplicationRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSupabaseApplicationRequestXLanguage) UnmarshalJSON(b []byte) error {
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
