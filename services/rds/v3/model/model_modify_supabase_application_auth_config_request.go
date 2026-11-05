package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ModifySupabaseApplicationAuthConfigRequest Request Object
type ModifySupabaseApplicationAuthConfigRequest struct {
	XLanguage *ModifySupabaseApplicationAuthConfigRequestXLanguage `json:"X-Language,omitempty"`

	ApplicationId string `json:"application_id"`

	Body *ModifySupabaseApplicationAuthConfigRequestBody `json:"body,omitempty"`
}

func (o ModifySupabaseApplicationAuthConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifySupabaseApplicationAuthConfigRequest struct{}"
	}

	return strings.Join([]string{"ModifySupabaseApplicationAuthConfigRequest", string(data)}, " ")
}

type ModifySupabaseApplicationAuthConfigRequestXLanguage struct {
	value string
}

type ModifySupabaseApplicationAuthConfigRequestXLanguageEnum struct {
	ZH_CN ModifySupabaseApplicationAuthConfigRequestXLanguage
	EN_US ModifySupabaseApplicationAuthConfigRequestXLanguage
}

func GetModifySupabaseApplicationAuthConfigRequestXLanguageEnum() ModifySupabaseApplicationAuthConfigRequestXLanguageEnum {
	return ModifySupabaseApplicationAuthConfigRequestXLanguageEnum{
		ZH_CN: ModifySupabaseApplicationAuthConfigRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ModifySupabaseApplicationAuthConfigRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ModifySupabaseApplicationAuthConfigRequestXLanguage) Value() string {
	return c.value
}

func (c ModifySupabaseApplicationAuthConfigRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifySupabaseApplicationAuthConfigRequestXLanguage) UnmarshalJSON(b []byte) error {
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
