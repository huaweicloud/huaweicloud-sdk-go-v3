package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ResetSupabaseApplicationApiKeyRequest Request Object
type ResetSupabaseApplicationApiKeyRequest struct {
	XLanguage *ResetSupabaseApplicationApiKeyRequestXLanguage `json:"X-Language,omitempty"`

	ApplicationId string `json:"application_id"`

	KeyId string `json:"key_id"`
}

func (o ResetSupabaseApplicationApiKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetSupabaseApplicationApiKeyRequest struct{}"
	}

	return strings.Join([]string{"ResetSupabaseApplicationApiKeyRequest", string(data)}, " ")
}

type ResetSupabaseApplicationApiKeyRequestXLanguage struct {
	value string
}

type ResetSupabaseApplicationApiKeyRequestXLanguageEnum struct {
	ZH_CN ResetSupabaseApplicationApiKeyRequestXLanguage
	EN_US ResetSupabaseApplicationApiKeyRequestXLanguage
}

func GetResetSupabaseApplicationApiKeyRequestXLanguageEnum() ResetSupabaseApplicationApiKeyRequestXLanguageEnum {
	return ResetSupabaseApplicationApiKeyRequestXLanguageEnum{
		ZH_CN: ResetSupabaseApplicationApiKeyRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ResetSupabaseApplicationApiKeyRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ResetSupabaseApplicationApiKeyRequestXLanguage) Value() string {
	return c.value
}

func (c ResetSupabaseApplicationApiKeyRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResetSupabaseApplicationApiKeyRequestXLanguage) UnmarshalJSON(b []byte) error {
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
