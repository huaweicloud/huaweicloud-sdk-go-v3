package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAgencyInfosRequest Request Object
type ListAgencyInfosRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *ListAgencyInfosRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 委托名称 **约束限制**: 不涉及。 **取值范围**: RDSAccessProjectResource **默认取值**: RDSAccessProjectResource
	AgencyName string `json:"agency_name"`
}

func (o ListAgencyInfosRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAgencyInfosRequest struct{}"
	}

	return strings.Join([]string{"ListAgencyInfosRequest", string(data)}, " ")
}

type ListAgencyInfosRequestXLanguage struct {
	value string
}

type ListAgencyInfosRequestXLanguageEnum struct {
	ZH_CN ListAgencyInfosRequestXLanguage
	EN_US ListAgencyInfosRequestXLanguage
}

func GetListAgencyInfosRequestXLanguageEnum() ListAgencyInfosRequestXLanguageEnum {
	return ListAgencyInfosRequestXLanguageEnum{
		ZH_CN: ListAgencyInfosRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListAgencyInfosRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListAgencyInfosRequestXLanguage) Value() string {
	return c.value
}

func (c ListAgencyInfosRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAgencyInfosRequestXLanguage) UnmarshalJSON(b []byte) error {
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
