package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateDryRunPolicyRequest Request Object
type UpdateDryRunPolicyRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// 策略的唯一标识符（ID）。
	PolicyId string `json:"policy_id"`

	// 选择接口返回的信息的语言
	XLanguage *UpdateDryRunPolicyRequestXLanguage `json:"X-Language,omitempty"`

	Body *UpdatePolicyReqBody `json:"body,omitempty"`
}

func (o UpdateDryRunPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDryRunPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateDryRunPolicyRequest", string(data)}, " ")
}

type UpdateDryRunPolicyRequestXLanguage struct {
	value string
}

type UpdateDryRunPolicyRequestXLanguageEnum struct {
	ZH_CN UpdateDryRunPolicyRequestXLanguage
	EN_US UpdateDryRunPolicyRequestXLanguage
}

func GetUpdateDryRunPolicyRequestXLanguageEnum() UpdateDryRunPolicyRequestXLanguageEnum {
	return UpdateDryRunPolicyRequestXLanguageEnum{
		ZH_CN: UpdateDryRunPolicyRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: UpdateDryRunPolicyRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c UpdateDryRunPolicyRequestXLanguage) Value() string {
	return c.value
}

func (c UpdateDryRunPolicyRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDryRunPolicyRequestXLanguage) UnmarshalJSON(b []byte) error {
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
