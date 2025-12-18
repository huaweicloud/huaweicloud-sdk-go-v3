package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateAgencyPolicyRequest Request Object
type UpdateAgencyPolicyRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *UpdateAgencyPolicyRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 委托名称 **约束限制**: 不涉及。 **取值范围**: RDSAccessProjectResource **默认取值**: RDSAccessProjectResource
	AgencyName string `json:"agency_name"`

	Body *UpdateAgencyPolicyRequestBody `json:"body,omitempty"`
}

func (o UpdateAgencyPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAgencyPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateAgencyPolicyRequest", string(data)}, " ")
}

type UpdateAgencyPolicyRequestXLanguage struct {
	value string
}

type UpdateAgencyPolicyRequestXLanguageEnum struct {
	ZH_CN UpdateAgencyPolicyRequestXLanguage
	EN_US UpdateAgencyPolicyRequestXLanguage
}

func GetUpdateAgencyPolicyRequestXLanguageEnum() UpdateAgencyPolicyRequestXLanguageEnum {
	return UpdateAgencyPolicyRequestXLanguageEnum{
		ZH_CN: UpdateAgencyPolicyRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: UpdateAgencyPolicyRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c UpdateAgencyPolicyRequestXLanguage) Value() string {
	return c.value
}

func (c UpdateAgencyPolicyRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAgencyPolicyRequestXLanguage) UnmarshalJSON(b []byte) error {
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
