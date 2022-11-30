package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ShowRecyclePolicyRequest struct {

	// 语言。默认值：en-us。
	XLanguage *ShowRecyclePolicyRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ShowRecyclePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRecyclePolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowRecyclePolicyRequest", string(data)}, " ")
}

type ShowRecyclePolicyRequestXLanguage struct {
	value string
}

type ShowRecyclePolicyRequestXLanguageEnum struct {
	ZH_CN ShowRecyclePolicyRequestXLanguage
	EN_US ShowRecyclePolicyRequestXLanguage
}

func GetShowRecyclePolicyRequestXLanguageEnum() ShowRecyclePolicyRequestXLanguageEnum {
	return ShowRecyclePolicyRequestXLanguageEnum{
		ZH_CN: ShowRecyclePolicyRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowRecyclePolicyRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowRecyclePolicyRequestXLanguage) Value() string {
	return c.value
}

func (c ShowRecyclePolicyRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowRecyclePolicyRequestXLanguage) UnmarshalJSON(b []byte) error {
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
