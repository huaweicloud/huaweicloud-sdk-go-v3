package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowConfigurationDetailRequest Request Object
type ShowConfigurationDetailRequest struct {

	// 语言,默认：en-us。
	XLanguage *ShowConfigurationDetailRequestXLanguage `json:"X-Language,omitempty"`

	// 参数模板ID
	ConfigId string `json:"config_id"`
}

func (o ShowConfigurationDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigurationDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowConfigurationDetailRequest", string(data)}, " ")
}

type ShowConfigurationDetailRequestXLanguage struct {
	value string
}

type ShowConfigurationDetailRequestXLanguageEnum struct {
	ZH_CN ShowConfigurationDetailRequestXLanguage
	EN_US ShowConfigurationDetailRequestXLanguage
}

func GetShowConfigurationDetailRequestXLanguageEnum() ShowConfigurationDetailRequestXLanguageEnum {
	return ShowConfigurationDetailRequestXLanguageEnum{
		ZH_CN: ShowConfigurationDetailRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowConfigurationDetailRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowConfigurationDetailRequestXLanguage) Value() string {
	return c.value
}

func (c ShowConfigurationDetailRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowConfigurationDetailRequestXLanguage) UnmarshalJSON(b []byte) error {
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
