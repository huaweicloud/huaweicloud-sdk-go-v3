package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowEpsRemainingQuotaRequest Request Object
type ShowEpsRemainingQuotaRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *ShowEpsRemainingQuotaRequestXLanguage `json:"X-Language,omitempty"`

	Body *ShowEpsRemainingQuotaRequestBody `json:"body,omitempty"`
}

func (o ShowEpsRemainingQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEpsRemainingQuotaRequest struct{}"
	}

	return strings.Join([]string{"ShowEpsRemainingQuotaRequest", string(data)}, " ")
}

type ShowEpsRemainingQuotaRequestXLanguage struct {
	value string
}

type ShowEpsRemainingQuotaRequestXLanguageEnum struct {
	ZH_CN ShowEpsRemainingQuotaRequestXLanguage
	EN_US ShowEpsRemainingQuotaRequestXLanguage
}

func GetShowEpsRemainingQuotaRequestXLanguageEnum() ShowEpsRemainingQuotaRequestXLanguageEnum {
	return ShowEpsRemainingQuotaRequestXLanguageEnum{
		ZH_CN: ShowEpsRemainingQuotaRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowEpsRemainingQuotaRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowEpsRemainingQuotaRequestXLanguage) Value() string {
	return c.value
}

func (c ShowEpsRemainingQuotaRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowEpsRemainingQuotaRequestXLanguage) UnmarshalJSON(b []byte) error {
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
