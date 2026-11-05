package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSupabaseApplicationAuthConfigRequest Request Object
type ShowSupabaseApplicationAuthConfigRequest struct {
	XLanguage *ShowSupabaseApplicationAuthConfigRequestXLanguage `json:"X-Language,omitempty"`

	ApplicationId string `json:"application_id"`
}

func (o ShowSupabaseApplicationAuthConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSupabaseApplicationAuthConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowSupabaseApplicationAuthConfigRequest", string(data)}, " ")
}

type ShowSupabaseApplicationAuthConfigRequestXLanguage struct {
	value string
}

type ShowSupabaseApplicationAuthConfigRequestXLanguageEnum struct {
	ZH_CN ShowSupabaseApplicationAuthConfigRequestXLanguage
	EN_US ShowSupabaseApplicationAuthConfigRequestXLanguage
}

func GetShowSupabaseApplicationAuthConfigRequestXLanguageEnum() ShowSupabaseApplicationAuthConfigRequestXLanguageEnum {
	return ShowSupabaseApplicationAuthConfigRequestXLanguageEnum{
		ZH_CN: ShowSupabaseApplicationAuthConfigRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowSupabaseApplicationAuthConfigRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowSupabaseApplicationAuthConfigRequestXLanguage) Value() string {
	return c.value
}

func (c ShowSupabaseApplicationAuthConfigRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSupabaseApplicationAuthConfigRequestXLanguage) UnmarshalJSON(b []byte) error {
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
