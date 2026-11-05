package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// RestartSupabaseApplicationRequest Request Object
type RestartSupabaseApplicationRequest struct {
	XLanguage *RestartSupabaseApplicationRequestXLanguage `json:"X-Language,omitempty"`

	ApplicationId string `json:"application_id"`
}

func (o RestartSupabaseApplicationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartSupabaseApplicationRequest struct{}"
	}

	return strings.Join([]string{"RestartSupabaseApplicationRequest", string(data)}, " ")
}

type RestartSupabaseApplicationRequestXLanguage struct {
	value string
}

type RestartSupabaseApplicationRequestXLanguageEnum struct {
	ZH_CN RestartSupabaseApplicationRequestXLanguage
	EN_US RestartSupabaseApplicationRequestXLanguage
}

func GetRestartSupabaseApplicationRequestXLanguageEnum() RestartSupabaseApplicationRequestXLanguageEnum {
	return RestartSupabaseApplicationRequestXLanguageEnum{
		ZH_CN: RestartSupabaseApplicationRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: RestartSupabaseApplicationRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c RestartSupabaseApplicationRequestXLanguage) Value() string {
	return c.value
}

func (c RestartSupabaseApplicationRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RestartSupabaseApplicationRequestXLanguage) UnmarshalJSON(b []byte) error {
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
