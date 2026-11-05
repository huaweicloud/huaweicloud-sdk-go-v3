package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SwitchSupabaseApplicationEipBindRequest Request Object
type SwitchSupabaseApplicationEipBindRequest struct {
	XLanguage *SwitchSupabaseApplicationEipBindRequestXLanguage `json:"X-Language,omitempty"`

	ApplicationId string `json:"application_id"`

	Body *SwitchSupabaseApplicationEipBindRequestBody `json:"body,omitempty"`
}

func (o SwitchSupabaseApplicationEipBindRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchSupabaseApplicationEipBindRequest struct{}"
	}

	return strings.Join([]string{"SwitchSupabaseApplicationEipBindRequest", string(data)}, " ")
}

type SwitchSupabaseApplicationEipBindRequestXLanguage struct {
	value string
}

type SwitchSupabaseApplicationEipBindRequestXLanguageEnum struct {
	ZH_CN SwitchSupabaseApplicationEipBindRequestXLanguage
	EN_US SwitchSupabaseApplicationEipBindRequestXLanguage
}

func GetSwitchSupabaseApplicationEipBindRequestXLanguageEnum() SwitchSupabaseApplicationEipBindRequestXLanguageEnum {
	return SwitchSupabaseApplicationEipBindRequestXLanguageEnum{
		ZH_CN: SwitchSupabaseApplicationEipBindRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: SwitchSupabaseApplicationEipBindRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c SwitchSupabaseApplicationEipBindRequestXLanguage) Value() string {
	return c.value
}

func (c SwitchSupabaseApplicationEipBindRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SwitchSupabaseApplicationEipBindRequestXLanguage) UnmarshalJSON(b []byte) error {
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
