package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SwitchSupabaseApplicationSslRequest Request Object
type SwitchSupabaseApplicationSslRequest struct {
	XLanguage *SwitchSupabaseApplicationSslRequestXLanguage `json:"X-Language,omitempty"`

	ApplicationId string `json:"application_id"`

	Body *SwitchSupabaseApplicationSslRequestBody `json:"body,omitempty"`
}

func (o SwitchSupabaseApplicationSslRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchSupabaseApplicationSslRequest struct{}"
	}

	return strings.Join([]string{"SwitchSupabaseApplicationSslRequest", string(data)}, " ")
}

type SwitchSupabaseApplicationSslRequestXLanguage struct {
	value string
}

type SwitchSupabaseApplicationSslRequestXLanguageEnum struct {
	ZH_CN SwitchSupabaseApplicationSslRequestXLanguage
	EN_US SwitchSupabaseApplicationSslRequestXLanguage
}

func GetSwitchSupabaseApplicationSslRequestXLanguageEnum() SwitchSupabaseApplicationSslRequestXLanguageEnum {
	return SwitchSupabaseApplicationSslRequestXLanguageEnum{
		ZH_CN: SwitchSupabaseApplicationSslRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: SwitchSupabaseApplicationSslRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c SwitchSupabaseApplicationSslRequestXLanguage) Value() string {
	return c.value
}

func (c SwitchSupabaseApplicationSslRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SwitchSupabaseApplicationSslRequestXLanguage) UnmarshalJSON(b []byte) error {
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
