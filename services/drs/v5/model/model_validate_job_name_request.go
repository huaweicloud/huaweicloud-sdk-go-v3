package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ValidateJobNameRequest Request Object
type ValidateJobNameRequest struct {

	// 请求语言类型。
	XLanguage *ValidateJobNameRequestXLanguage `json:"X-Language,omitempty"`

	Body *CheckJobNameReq `json:"body,omitempty"`
}

func (o ValidateJobNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateJobNameRequest struct{}"
	}

	return strings.Join([]string{"ValidateJobNameRequest", string(data)}, " ")
}

type ValidateJobNameRequestXLanguage struct {
	value string
}

type ValidateJobNameRequestXLanguageEnum struct {
	EN_US ValidateJobNameRequestXLanguage
	ZH_CN ValidateJobNameRequestXLanguage
}

func GetValidateJobNameRequestXLanguageEnum() ValidateJobNameRequestXLanguageEnum {
	return ValidateJobNameRequestXLanguageEnum{
		EN_US: ValidateJobNameRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ValidateJobNameRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ValidateJobNameRequestXLanguage) Value() string {
	return c.value
}

func (c ValidateJobNameRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ValidateJobNameRequestXLanguage) UnmarshalJSON(b []byte) error {
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
