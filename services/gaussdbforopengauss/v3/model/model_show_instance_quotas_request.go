package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowInstanceQuotasRequest Request Object
type ShowInstanceQuotasRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn - en-us  **默认取值**: en-us
	XLanguage *ShowInstanceQuotasRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ShowInstanceQuotasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceQuotasRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceQuotasRequest", string(data)}, " ")
}

type ShowInstanceQuotasRequestXLanguage struct {
	value string
}

type ShowInstanceQuotasRequestXLanguageEnum struct {
	ZH_CN ShowInstanceQuotasRequestXLanguage
	EN_US ShowInstanceQuotasRequestXLanguage
}

func GetShowInstanceQuotasRequestXLanguageEnum() ShowInstanceQuotasRequestXLanguageEnum {
	return ShowInstanceQuotasRequestXLanguageEnum{
		ZH_CN: ShowInstanceQuotasRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowInstanceQuotasRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowInstanceQuotasRequestXLanguage) Value() string {
	return c.value
}

func (c ShowInstanceQuotasRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstanceQuotasRequestXLanguage) UnmarshalJSON(b []byte) error {
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
