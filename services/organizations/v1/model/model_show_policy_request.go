package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowPolicyRequest Request Object
type ShowPolicyRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// 策略的唯一标识符（ID）。
	PolicyId string `json:"policy_id"`

	// 选择接口返回的信息的语言
	XLanguage *ShowPolicyRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ShowPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowPolicyRequest", string(data)}, " ")
}

type ShowPolicyRequestXLanguage struct {
	value string
}

type ShowPolicyRequestXLanguageEnum struct {
	ZH_CN ShowPolicyRequestXLanguage
	EN_US ShowPolicyRequestXLanguage
}

func GetShowPolicyRequestXLanguageEnum() ShowPolicyRequestXLanguageEnum {
	return ShowPolicyRequestXLanguageEnum{
		ZH_CN: ShowPolicyRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowPolicyRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowPolicyRequestXLanguage) Value() string {
	return c.value
}

func (c ShowPolicyRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowPolicyRequestXLanguage) UnmarshalJSON(b []byte) error {
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
