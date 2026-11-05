package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StopSupabaseApplicationRequest Request Object
type StopSupabaseApplicationRequest struct {
	XLanguage *StopSupabaseApplicationRequestXLanguage `json:"X-Language,omitempty"`

	ApplicationId string `json:"application_id"`
}

func (o StopSupabaseApplicationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopSupabaseApplicationRequest struct{}"
	}

	return strings.Join([]string{"StopSupabaseApplicationRequest", string(data)}, " ")
}

type StopSupabaseApplicationRequestXLanguage struct {
	value string
}

type StopSupabaseApplicationRequestXLanguageEnum struct {
	ZH_CN StopSupabaseApplicationRequestXLanguage
	EN_US StopSupabaseApplicationRequestXLanguage
}

func GetStopSupabaseApplicationRequestXLanguageEnum() StopSupabaseApplicationRequestXLanguageEnum {
	return StopSupabaseApplicationRequestXLanguageEnum{
		ZH_CN: StopSupabaseApplicationRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: StopSupabaseApplicationRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c StopSupabaseApplicationRequestXLanguage) Value() string {
	return c.value
}

func (c StopSupabaseApplicationRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StopSupabaseApplicationRequestXLanguage) UnmarshalJSON(b []byte) error {
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
