package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSupabaseApplicationApiKeyValueRequest Request Object
type ShowSupabaseApplicationApiKeyValueRequest struct {
	XLanguage *ShowSupabaseApplicationApiKeyValueRequestXLanguage `json:"X-Language,omitempty"`

	ApplicationId string `json:"application_id"`

	// ApiKey ID。
	KeyId string `json:"key_id"`
}

func (o ShowSupabaseApplicationApiKeyValueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSupabaseApplicationApiKeyValueRequest struct{}"
	}

	return strings.Join([]string{"ShowSupabaseApplicationApiKeyValueRequest", string(data)}, " ")
}

type ShowSupabaseApplicationApiKeyValueRequestXLanguage struct {
	value string
}

type ShowSupabaseApplicationApiKeyValueRequestXLanguageEnum struct {
	ZH_CN ShowSupabaseApplicationApiKeyValueRequestXLanguage
	EN_US ShowSupabaseApplicationApiKeyValueRequestXLanguage
}

func GetShowSupabaseApplicationApiKeyValueRequestXLanguageEnum() ShowSupabaseApplicationApiKeyValueRequestXLanguageEnum {
	return ShowSupabaseApplicationApiKeyValueRequestXLanguageEnum{
		ZH_CN: ShowSupabaseApplicationApiKeyValueRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowSupabaseApplicationApiKeyValueRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowSupabaseApplicationApiKeyValueRequestXLanguage) Value() string {
	return c.value
}

func (c ShowSupabaseApplicationApiKeyValueRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSupabaseApplicationApiKeyValueRequestXLanguage) UnmarshalJSON(b []byte) error {
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
