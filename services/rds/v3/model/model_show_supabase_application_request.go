package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSupabaseApplicationRequest Request Object
type ShowSupabaseApplicationRequest struct {
	XLanguage *ShowSupabaseApplicationRequestXLanguage `json:"X-Language,omitempty"`

	// Supabase应用ID。
	ApplicationId string `json:"application_id"`
}

func (o ShowSupabaseApplicationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSupabaseApplicationRequest struct{}"
	}

	return strings.Join([]string{"ShowSupabaseApplicationRequest", string(data)}, " ")
}

type ShowSupabaseApplicationRequestXLanguage struct {
	value string
}

type ShowSupabaseApplicationRequestXLanguageEnum struct {
	ZH_CN ShowSupabaseApplicationRequestXLanguage
	EN_US ShowSupabaseApplicationRequestXLanguage
}

func GetShowSupabaseApplicationRequestXLanguageEnum() ShowSupabaseApplicationRequestXLanguageEnum {
	return ShowSupabaseApplicationRequestXLanguageEnum{
		ZH_CN: ShowSupabaseApplicationRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowSupabaseApplicationRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowSupabaseApplicationRequestXLanguage) Value() string {
	return c.value
}

func (c ShowSupabaseApplicationRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSupabaseApplicationRequestXLanguage) UnmarshalJSON(b []byte) error {
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
