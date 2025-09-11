package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SearchAutoEnlargePolicyRequest Request Object
type SearchAutoEnlargePolicyRequest struct {

	// 语言
	XLanguage *SearchAutoEnlargePolicyRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`
}

func (o SearchAutoEnlargePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchAutoEnlargePolicyRequest struct{}"
	}

	return strings.Join([]string{"SearchAutoEnlargePolicyRequest", string(data)}, " ")
}

type SearchAutoEnlargePolicyRequestXLanguage struct {
	value string
}

type SearchAutoEnlargePolicyRequestXLanguageEnum struct {
	ZH_CN SearchAutoEnlargePolicyRequestXLanguage
	EN_US SearchAutoEnlargePolicyRequestXLanguage
}

func GetSearchAutoEnlargePolicyRequestXLanguageEnum() SearchAutoEnlargePolicyRequestXLanguageEnum {
	return SearchAutoEnlargePolicyRequestXLanguageEnum{
		ZH_CN: SearchAutoEnlargePolicyRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: SearchAutoEnlargePolicyRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c SearchAutoEnlargePolicyRequestXLanguage) Value() string {
	return c.value
}

func (c SearchAutoEnlargePolicyRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SearchAutoEnlargePolicyRequestXLanguage) UnmarshalJSON(b []byte) error {
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
