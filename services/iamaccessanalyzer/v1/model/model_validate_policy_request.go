package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ValidatePolicyRequest Request Object
type ValidatePolicyRequest struct {

	// 单页最大结果数。
	Limit *int32 `json:"limit,omitempty"`

	// 页面标记。
	Marker *string `json:"marker,omitempty"`

	// 返回消息的语言，默认值为'zh-cn'。 - zh-cn：中文 - en-us：英文
	XLanguage *ValidatePolicyRequestXLanguage `json:"X-Language,omitempty"`

	Body *ValidatePolicyReqBody `json:"body,omitempty"`
}

func (o ValidatePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidatePolicyRequest struct{}"
	}

	return strings.Join([]string{"ValidatePolicyRequest", string(data)}, " ")
}

type ValidatePolicyRequestXLanguage struct {
	value string
}

type ValidatePolicyRequestXLanguageEnum struct {
	ZH_CN ValidatePolicyRequestXLanguage
	EN_US ValidatePolicyRequestXLanguage
}

func GetValidatePolicyRequestXLanguageEnum() ValidatePolicyRequestXLanguageEnum {
	return ValidatePolicyRequestXLanguageEnum{
		ZH_CN: ValidatePolicyRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ValidatePolicyRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ValidatePolicyRequestXLanguage) Value() string {
	return c.value
}

func (c ValidatePolicyRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ValidatePolicyRequestXLanguage) UnmarshalJSON(b []byte) error {
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
