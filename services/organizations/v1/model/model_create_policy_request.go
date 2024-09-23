package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreatePolicyRequest Request Object
type CreatePolicyRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// 选择接口返回的信息的语言
	XLanguage *CreatePolicyRequestXLanguage `json:"X-Language,omitempty"`

	Body *CreatePolicyReqBody `json:"body,omitempty"`
}

func (o CreatePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePolicyRequest struct{}"
	}

	return strings.Join([]string{"CreatePolicyRequest", string(data)}, " ")
}

type CreatePolicyRequestXLanguage struct {
	value string
}

type CreatePolicyRequestXLanguageEnum struct {
	ZH_CN CreatePolicyRequestXLanguage
	EN_US CreatePolicyRequestXLanguage
}

func GetCreatePolicyRequestXLanguageEnum() CreatePolicyRequestXLanguageEnum {
	return CreatePolicyRequestXLanguageEnum{
		ZH_CN: CreatePolicyRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: CreatePolicyRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c CreatePolicyRequestXLanguage) Value() string {
	return c.value
}

func (c CreatePolicyRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePolicyRequestXLanguage) UnmarshalJSON(b []byte) error {
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
