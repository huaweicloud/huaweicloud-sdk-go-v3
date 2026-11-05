package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ModifySupabaseApplicationPasswordRequest Request Object
type ModifySupabaseApplicationPasswordRequest struct {
	XLanguage *ModifySupabaseApplicationPasswordRequestXLanguage `json:"X-Language,omitempty"`

	ApplicationId string `json:"application_id"`

	Body *ModifySupabaseApplicationPasswordRequestBody `json:"body,omitempty"`
}

func (o ModifySupabaseApplicationPasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifySupabaseApplicationPasswordRequest struct{}"
	}

	return strings.Join([]string{"ModifySupabaseApplicationPasswordRequest", string(data)}, " ")
}

type ModifySupabaseApplicationPasswordRequestXLanguage struct {
	value string
}

type ModifySupabaseApplicationPasswordRequestXLanguageEnum struct {
	ZH_CN ModifySupabaseApplicationPasswordRequestXLanguage
	EN_US ModifySupabaseApplicationPasswordRequestXLanguage
}

func GetModifySupabaseApplicationPasswordRequestXLanguageEnum() ModifySupabaseApplicationPasswordRequestXLanguageEnum {
	return ModifySupabaseApplicationPasswordRequestXLanguageEnum{
		ZH_CN: ModifySupabaseApplicationPasswordRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ModifySupabaseApplicationPasswordRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ModifySupabaseApplicationPasswordRequestXLanguage) Value() string {
	return c.value
}

func (c ModifySupabaseApplicationPasswordRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifySupabaseApplicationPasswordRequestXLanguage) UnmarshalJSON(b []byte) error {
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
