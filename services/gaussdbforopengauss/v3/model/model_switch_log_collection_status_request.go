package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SwitchLogCollectionStatusRequest Request Object
type SwitchLogCollectionStatusRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *SwitchLogCollectionStatusRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**: 采集状态。 **约束限制**: 不涉及。 **取值范围**: - ON：开始采集。 - OFF：关闭采集。 **默认取值**: 不涉及。
	Status string `json:"status"`
}

func (o SwitchLogCollectionStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchLogCollectionStatusRequest struct{}"
	}

	return strings.Join([]string{"SwitchLogCollectionStatusRequest", string(data)}, " ")
}

type SwitchLogCollectionStatusRequestXLanguage struct {
	value string
}

type SwitchLogCollectionStatusRequestXLanguageEnum struct {
	ZH_CN SwitchLogCollectionStatusRequestXLanguage
	EN_US SwitchLogCollectionStatusRequestXLanguage
}

func GetSwitchLogCollectionStatusRequestXLanguageEnum() SwitchLogCollectionStatusRequestXLanguageEnum {
	return SwitchLogCollectionStatusRequestXLanguageEnum{
		ZH_CN: SwitchLogCollectionStatusRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: SwitchLogCollectionStatusRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c SwitchLogCollectionStatusRequestXLanguage) Value() string {
	return c.value
}

func (c SwitchLogCollectionStatusRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SwitchLogCollectionStatusRequestXLanguage) UnmarshalJSON(b []byte) error {
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
