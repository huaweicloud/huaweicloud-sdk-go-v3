package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateDryRunPolicyRequest Request Object
type CreateDryRunPolicyRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// 选择接口返回的信息的语言
	XLanguage *CreateDryRunPolicyRequestXLanguage `json:"X-Language,omitempty"`

	Body *CreateDryRunPolicyReqBody `json:"body,omitempty"`
}

func (o CreateDryRunPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDryRunPolicyRequest struct{}"
	}

	return strings.Join([]string{"CreateDryRunPolicyRequest", string(data)}, " ")
}

type CreateDryRunPolicyRequestXLanguage struct {
	value string
}

type CreateDryRunPolicyRequestXLanguageEnum struct {
	ZH_CN CreateDryRunPolicyRequestXLanguage
	EN_US CreateDryRunPolicyRequestXLanguage
}

func GetCreateDryRunPolicyRequestXLanguageEnum() CreateDryRunPolicyRequestXLanguageEnum {
	return CreateDryRunPolicyRequestXLanguageEnum{
		ZH_CN: CreateDryRunPolicyRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: CreateDryRunPolicyRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c CreateDryRunPolicyRequestXLanguage) Value() string {
	return c.value
}

func (c CreateDryRunPolicyRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDryRunPolicyRequestXLanguage) UnmarshalJSON(b []byte) error {
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
