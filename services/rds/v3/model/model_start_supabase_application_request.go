package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StartSupabaseApplicationRequest Request Object
type StartSupabaseApplicationRequest struct {
	XLanguage *StartSupabaseApplicationRequestXLanguage `json:"X-Language,omitempty"`

	ApplicationId string `json:"application_id"`
}

func (o StartSupabaseApplicationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartSupabaseApplicationRequest struct{}"
	}

	return strings.Join([]string{"StartSupabaseApplicationRequest", string(data)}, " ")
}

type StartSupabaseApplicationRequestXLanguage struct {
	value string
}

type StartSupabaseApplicationRequestXLanguageEnum struct {
	ZH_CN StartSupabaseApplicationRequestXLanguage
	EN_US StartSupabaseApplicationRequestXLanguage
}

func GetStartSupabaseApplicationRequestXLanguageEnum() StartSupabaseApplicationRequestXLanguageEnum {
	return StartSupabaseApplicationRequestXLanguageEnum{
		ZH_CN: StartSupabaseApplicationRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: StartSupabaseApplicationRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c StartSupabaseApplicationRequestXLanguage) Value() string {
	return c.value
}

func (c StartSupabaseApplicationRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StartSupabaseApplicationRequestXLanguage) UnmarshalJSON(b []byte) error {
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
