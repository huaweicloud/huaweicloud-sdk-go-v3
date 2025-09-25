package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowExpansionParametersRequest Request Object
type ShowExpansionParametersRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *ShowExpansionParametersRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ShowExpansionParametersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowExpansionParametersRequest struct{}"
	}

	return strings.Join([]string{"ShowExpansionParametersRequest", string(data)}, " ")
}

type ShowExpansionParametersRequestXLanguage struct {
	value string
}

type ShowExpansionParametersRequestXLanguageEnum struct {
	ZH_CN ShowExpansionParametersRequestXLanguage
	EN_US ShowExpansionParametersRequestXLanguage
}

func GetShowExpansionParametersRequestXLanguageEnum() ShowExpansionParametersRequestXLanguageEnum {
	return ShowExpansionParametersRequestXLanguageEnum{
		ZH_CN: ShowExpansionParametersRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowExpansionParametersRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowExpansionParametersRequestXLanguage) Value() string {
	return c.value
}

func (c ShowExpansionParametersRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowExpansionParametersRequestXLanguage) UnmarshalJSON(b []byte) error {
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
