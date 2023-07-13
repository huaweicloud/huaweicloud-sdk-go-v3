package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ValidateWeakPasswordRequest Request Object
type ValidateWeakPasswordRequest struct {

	// 语言。默认值：en-us。
	XLanguage *ValidateWeakPasswordRequestXLanguage `json:"X-Language,omitempty"`

	Body *WeakPasswordRequestBody `json:"body,omitempty"`
}

func (o ValidateWeakPasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateWeakPasswordRequest struct{}"
	}

	return strings.Join([]string{"ValidateWeakPasswordRequest", string(data)}, " ")
}

type ValidateWeakPasswordRequestXLanguage struct {
	value string
}

type ValidateWeakPasswordRequestXLanguageEnum struct {
	ZH_CN ValidateWeakPasswordRequestXLanguage
	EN_US ValidateWeakPasswordRequestXLanguage
}

func GetValidateWeakPasswordRequestXLanguageEnum() ValidateWeakPasswordRequestXLanguageEnum {
	return ValidateWeakPasswordRequestXLanguageEnum{
		ZH_CN: ValidateWeakPasswordRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ValidateWeakPasswordRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ValidateWeakPasswordRequestXLanguage) Value() string {
	return c.value
}

func (c ValidateWeakPasswordRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ValidateWeakPasswordRequestXLanguage) UnmarshalJSON(b []byte) error {
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
