package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CheckWeakpwdRequest Request Object
type CheckWeakpwdRequest struct {

	// 语言。
	XLanguage *CheckWeakpwdRequestXLanguage `json:"X-Language,omitempty"`

	Body *CheckWeakPasswordRequest `json:"body,omitempty"`
}

func (o CheckWeakpwdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckWeakpwdRequest struct{}"
	}

	return strings.Join([]string{"CheckWeakpwdRequest", string(data)}, " ")
}

type CheckWeakpwdRequestXLanguage struct {
	value string
}

type CheckWeakpwdRequestXLanguageEnum struct {
	ZH_CN CheckWeakpwdRequestXLanguage
	EN_US CheckWeakpwdRequestXLanguage
}

func GetCheckWeakpwdRequestXLanguageEnum() CheckWeakpwdRequestXLanguageEnum {
	return CheckWeakpwdRequestXLanguageEnum{
		ZH_CN: CheckWeakpwdRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: CheckWeakpwdRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c CheckWeakpwdRequestXLanguage) Value() string {
	return c.value
}

func (c CheckWeakpwdRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CheckWeakpwdRequestXLanguage) UnmarshalJSON(b []byte) error {
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
