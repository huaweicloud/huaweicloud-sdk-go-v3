package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ModifySupabaseApplicationStorageRequest Request Object
type ModifySupabaseApplicationStorageRequest struct {
	XLanguage *ModifySupabaseApplicationStorageRequestXLanguage `json:"X-Language,omitempty"`

	ApplicationId string `json:"application_id"`

	Body *ModifySupabaseApplicationStorageRequestBody `json:"body,omitempty"`
}

func (o ModifySupabaseApplicationStorageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifySupabaseApplicationStorageRequest struct{}"
	}

	return strings.Join([]string{"ModifySupabaseApplicationStorageRequest", string(data)}, " ")
}

type ModifySupabaseApplicationStorageRequestXLanguage struct {
	value string
}

type ModifySupabaseApplicationStorageRequestXLanguageEnum struct {
	ZH_CN ModifySupabaseApplicationStorageRequestXLanguage
	EN_US ModifySupabaseApplicationStorageRequestXLanguage
}

func GetModifySupabaseApplicationStorageRequestXLanguageEnum() ModifySupabaseApplicationStorageRequestXLanguageEnum {
	return ModifySupabaseApplicationStorageRequestXLanguageEnum{
		ZH_CN: ModifySupabaseApplicationStorageRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ModifySupabaseApplicationStorageRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ModifySupabaseApplicationStorageRequestXLanguage) Value() string {
	return c.value
}

func (c ModifySupabaseApplicationStorageRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifySupabaseApplicationStorageRequestXLanguage) UnmarshalJSON(b []byte) error {
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
