package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateSecurityPolicyControlRequest Request Object
type UpdateSecurityPolicyControlRequest struct {

	// 语言，用于国际化。 - en-us：英文 - zh-cn：中文
	XLanguage *UpdateSecurityPolicyControlRequestXLanguage `json:"X-Language,omitempty"`

	Body *UpdateSecurityPolicyControlReq `json:"body,omitempty"`
}

func (o UpdateSecurityPolicyControlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityPolicyControlRequest struct{}"
	}

	return strings.Join([]string{"UpdateSecurityPolicyControlRequest", string(data)}, " ")
}

type UpdateSecurityPolicyControlRequestXLanguage struct {
	value string
}

type UpdateSecurityPolicyControlRequestXLanguageEnum struct {
	EN_US UpdateSecurityPolicyControlRequestXLanguage
	ZH_CN UpdateSecurityPolicyControlRequestXLanguage
}

func GetUpdateSecurityPolicyControlRequestXLanguageEnum() UpdateSecurityPolicyControlRequestXLanguageEnum {
	return UpdateSecurityPolicyControlRequestXLanguageEnum{
		EN_US: UpdateSecurityPolicyControlRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: UpdateSecurityPolicyControlRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c UpdateSecurityPolicyControlRequestXLanguage) Value() string {
	return c.value
}

func (c UpdateSecurityPolicyControlRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateSecurityPolicyControlRequestXLanguage) UnmarshalJSON(b []byte) error {
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
