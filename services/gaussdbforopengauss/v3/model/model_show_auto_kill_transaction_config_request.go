package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowAutoKillTransactionConfigRequest Request Object
type ShowAutoKillTransactionConfigRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *ShowAutoKillTransactionConfigRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**: 自动查杀配置类型。 **约束限制**: 不涉及。 **取值范围**: - exec_time：代表长事务。 - xlog_quantity：代表大事务。  **默认取值**: 不涉及。
	Type string `json:"type"`
}

func (o ShowAutoKillTransactionConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoKillTransactionConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowAutoKillTransactionConfigRequest", string(data)}, " ")
}

type ShowAutoKillTransactionConfigRequestXLanguage struct {
	value string
}

type ShowAutoKillTransactionConfigRequestXLanguageEnum struct {
	ZH_CN ShowAutoKillTransactionConfigRequestXLanguage
	EN_US ShowAutoKillTransactionConfigRequestXLanguage
}

func GetShowAutoKillTransactionConfigRequestXLanguageEnum() ShowAutoKillTransactionConfigRequestXLanguageEnum {
	return ShowAutoKillTransactionConfigRequestXLanguageEnum{
		ZH_CN: ShowAutoKillTransactionConfigRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowAutoKillTransactionConfigRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowAutoKillTransactionConfigRequestXLanguage) Value() string {
	return c.value
}

func (c ShowAutoKillTransactionConfigRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAutoKillTransactionConfigRequestXLanguage) UnmarshalJSON(b []byte) error {
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
