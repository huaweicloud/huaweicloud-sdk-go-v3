package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteSupabaseApplicationRequest Request Object
type DeleteSupabaseApplicationRequest struct {
	XLanguage *DeleteSupabaseApplicationRequestXLanguage `json:"X-Language,omitempty"`

	ApplicationId string `json:"application_id"`
}

func (o DeleteSupabaseApplicationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSupabaseApplicationRequest struct{}"
	}

	return strings.Join([]string{"DeleteSupabaseApplicationRequest", string(data)}, " ")
}

type DeleteSupabaseApplicationRequestXLanguage struct {
	value string
}

type DeleteSupabaseApplicationRequestXLanguageEnum struct {
	ZH_CN DeleteSupabaseApplicationRequestXLanguage
	EN_US DeleteSupabaseApplicationRequestXLanguage
}

func GetDeleteSupabaseApplicationRequestXLanguageEnum() DeleteSupabaseApplicationRequestXLanguageEnum {
	return DeleteSupabaseApplicationRequestXLanguageEnum{
		ZH_CN: DeleteSupabaseApplicationRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: DeleteSupabaseApplicationRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c DeleteSupabaseApplicationRequestXLanguage) Value() string {
	return c.value
}

func (c DeleteSupabaseApplicationRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteSupabaseApplicationRequestXLanguage) UnmarshalJSON(b []byte) error {
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
