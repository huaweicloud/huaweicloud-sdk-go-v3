package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowRedistributionParametersRequest Request Object
type ShowRedistributionParametersRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *ShowRedistributionParametersRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ShowRedistributionParametersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRedistributionParametersRequest struct{}"
	}

	return strings.Join([]string{"ShowRedistributionParametersRequest", string(data)}, " ")
}

type ShowRedistributionParametersRequestXLanguage struct {
	value string
}

type ShowRedistributionParametersRequestXLanguageEnum struct {
	ZH_CN ShowRedistributionParametersRequestXLanguage
	EN_US ShowRedistributionParametersRequestXLanguage
}

func GetShowRedistributionParametersRequestXLanguageEnum() ShowRedistributionParametersRequestXLanguageEnum {
	return ShowRedistributionParametersRequestXLanguageEnum{
		ZH_CN: ShowRedistributionParametersRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowRedistributionParametersRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowRedistributionParametersRequestXLanguage) Value() string {
	return c.value
}

func (c ShowRedistributionParametersRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowRedistributionParametersRequestXLanguage) UnmarshalJSON(b []byte) error {
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
