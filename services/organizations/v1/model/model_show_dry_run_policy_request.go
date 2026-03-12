package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowDryRunPolicyRequest Request Object
type ShowDryRunPolicyRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// 策略的唯一标识符（ID）。
	PolicyId string `json:"policy_id"`

	// 选择接口返回的信息的语言
	XLanguage *ShowDryRunPolicyRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ShowDryRunPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDryRunPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowDryRunPolicyRequest", string(data)}, " ")
}

type ShowDryRunPolicyRequestXLanguage struct {
	value string
}

type ShowDryRunPolicyRequestXLanguageEnum struct {
	ZH_CN ShowDryRunPolicyRequestXLanguage
	EN_US ShowDryRunPolicyRequestXLanguage
}

func GetShowDryRunPolicyRequestXLanguageEnum() ShowDryRunPolicyRequestXLanguageEnum {
	return ShowDryRunPolicyRequestXLanguageEnum{
		ZH_CN: ShowDryRunPolicyRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowDryRunPolicyRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowDryRunPolicyRequestXLanguage) Value() string {
	return c.value
}

func (c ShowDryRunPolicyRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDryRunPolicyRequestXLanguage) UnmarshalJSON(b []byte) error {
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
