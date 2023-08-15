package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ResetConfigurationRequest Request Object
type ResetConfigurationRequest struct {

	// 需重置的参数模板ID。
	ConfigId string `json:"config_id"`

	// 语言。
	XLanguage *ResetConfigurationRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ResetConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetConfigurationRequest struct{}"
	}

	return strings.Join([]string{"ResetConfigurationRequest", string(data)}, " ")
}

type ResetConfigurationRequestXLanguage struct {
	value string
}

type ResetConfigurationRequestXLanguageEnum struct {
	ZH_CN ResetConfigurationRequestXLanguage
	EN_US ResetConfigurationRequestXLanguage
}

func GetResetConfigurationRequestXLanguageEnum() ResetConfigurationRequestXLanguageEnum {
	return ResetConfigurationRequestXLanguageEnum{
		ZH_CN: ResetConfigurationRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ResetConfigurationRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ResetConfigurationRequestXLanguage) Value() string {
	return c.value
}

func (c ResetConfigurationRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResetConfigurationRequestXLanguage) UnmarshalJSON(b []byte) error {
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
