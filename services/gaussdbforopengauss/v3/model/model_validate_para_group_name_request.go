package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ValidateParaGroupNameRequest struct {

	// 语言
	XLanguage *ValidateParaGroupNameRequestXLanguage `json:"X-Language,omitempty"`

	// 参数组名称。
	Name string `json:"name"`
}

func (o ValidateParaGroupNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateParaGroupNameRequest struct{}"
	}

	return strings.Join([]string{"ValidateParaGroupNameRequest", string(data)}, " ")
}

type ValidateParaGroupNameRequestXLanguage struct {
	value string
}

type ValidateParaGroupNameRequestXLanguageEnum struct {
	ZH_CN ValidateParaGroupNameRequestXLanguage
	EN_US ValidateParaGroupNameRequestXLanguage
}

func GetValidateParaGroupNameRequestXLanguageEnum() ValidateParaGroupNameRequestXLanguageEnum {
	return ValidateParaGroupNameRequestXLanguageEnum{
		ZH_CN: ValidateParaGroupNameRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ValidateParaGroupNameRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ValidateParaGroupNameRequestXLanguage) Value() string {
	return c.value
}

func (c ValidateParaGroupNameRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ValidateParaGroupNameRequestXLanguage) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
