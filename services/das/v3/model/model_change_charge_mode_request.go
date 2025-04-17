package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ChangeChargeModeRequest Request Object
type ChangeChargeModeRequest struct {

	// 语言
	XLanguage *ChangeChargeModeRequestXLanguage `json:"X-Language,omitempty"`

	Body *ChangeChargeModeBody `json:"body,omitempty"`
}

func (o ChangeChargeModeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeChargeModeRequest struct{}"
	}

	return strings.Join([]string{"ChangeChargeModeRequest", string(data)}, " ")
}

type ChangeChargeModeRequestXLanguage struct {
	value string
}

type ChangeChargeModeRequestXLanguageEnum struct {
	ZH_CN ChangeChargeModeRequestXLanguage
	EN_US ChangeChargeModeRequestXLanguage
}

func GetChangeChargeModeRequestXLanguageEnum() ChangeChargeModeRequestXLanguageEnum {
	return ChangeChargeModeRequestXLanguageEnum{
		ZH_CN: ChangeChargeModeRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ChangeChargeModeRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ChangeChargeModeRequestXLanguage) Value() string {
	return c.value
}

func (c ChangeChargeModeRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChangeChargeModeRequestXLanguage) UnmarshalJSON(b []byte) error {
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
