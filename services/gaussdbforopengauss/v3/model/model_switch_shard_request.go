package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SwitchShardRequest Request Object
type SwitchShardRequest struct {

	// 语言
	XLanguage *SwitchShardRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	Body *SwitchShardRequestBody `json:"body,omitempty"`
}

func (o SwitchShardRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchShardRequest struct{}"
	}

	return strings.Join([]string{"SwitchShardRequest", string(data)}, " ")
}

type SwitchShardRequestXLanguage struct {
	value string
}

type SwitchShardRequestXLanguageEnum struct {
	ZH_CN SwitchShardRequestXLanguage
	EN_US SwitchShardRequestXLanguage
}

func GetSwitchShardRequestXLanguageEnum() SwitchShardRequestXLanguageEnum {
	return SwitchShardRequestXLanguageEnum{
		ZH_CN: SwitchShardRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: SwitchShardRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c SwitchShardRequestXLanguage) Value() string {
	return c.value
}

func (c SwitchShardRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SwitchShardRequestXLanguage) UnmarshalJSON(b []byte) error {
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
