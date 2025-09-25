package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ModifyAutoEnlargePolicyRequest Request Object
type ModifyAutoEnlargePolicyRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**:   - zh-cn   - en-us  **默认取值**: en-us
	XLanguage *ModifyAutoEnlargePolicyRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	Body *ModifyAutoEnlargePolicyRequestBody `json:"body,omitempty"`
}

func (o ModifyAutoEnlargePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyAutoEnlargePolicyRequest struct{}"
	}

	return strings.Join([]string{"ModifyAutoEnlargePolicyRequest", string(data)}, " ")
}

type ModifyAutoEnlargePolicyRequestXLanguage struct {
	value string
}

type ModifyAutoEnlargePolicyRequestXLanguageEnum struct {
	ZH_CN ModifyAutoEnlargePolicyRequestXLanguage
	EN_US ModifyAutoEnlargePolicyRequestXLanguage
}

func GetModifyAutoEnlargePolicyRequestXLanguageEnum() ModifyAutoEnlargePolicyRequestXLanguageEnum {
	return ModifyAutoEnlargePolicyRequestXLanguageEnum{
		ZH_CN: ModifyAutoEnlargePolicyRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ModifyAutoEnlargePolicyRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ModifyAutoEnlargePolicyRequestXLanguage) Value() string {
	return c.value
}

func (c ModifyAutoEnlargePolicyRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifyAutoEnlargePolicyRequestXLanguage) UnmarshalJSON(b []byte) error {
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
