package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type UpdatePolicyRequest struct {

	// 根、组织单元或帐号的唯一标识符（ID）。
	PolicyId string `json:"policy_id"`

	// 选择接口返回的信息的语言，默认为\"zh-cn\"中文，zh-cn中文，en-us英文
	XLanguage *UpdatePolicyRequestXLanguage `json:"X-Language,omitempty"`

	Body *UpdatePolicyReqBody `json:"body,omitempty"`
}

func (o UpdatePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdatePolicyRequest", string(data)}, " ")
}

type UpdatePolicyRequestXLanguage struct {
	value string
}

type UpdatePolicyRequestXLanguageEnum struct {
	ZH_CN UpdatePolicyRequestXLanguage
	EN_US UpdatePolicyRequestXLanguage
}

func GetUpdatePolicyRequestXLanguageEnum() UpdatePolicyRequestXLanguageEnum {
	return UpdatePolicyRequestXLanguageEnum{
		ZH_CN: UpdatePolicyRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: UpdatePolicyRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c UpdatePolicyRequestXLanguage) Value() string {
	return c.value
}

func (c UpdatePolicyRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdatePolicyRequestXLanguage) UnmarshalJSON(b []byte) error {
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
