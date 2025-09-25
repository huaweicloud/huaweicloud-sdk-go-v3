package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ModifyHotfixesRequest Request Object
type ModifyHotfixesRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *ModifyHotfixesRequestXLanguage `json:"X-Language,omitempty"`

	Body *ModifyHotfixesRequestBody `json:"body,omitempty"`
}

func (o ModifyHotfixesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyHotfixesRequest struct{}"
	}

	return strings.Join([]string{"ModifyHotfixesRequest", string(data)}, " ")
}

type ModifyHotfixesRequestXLanguage struct {
	value string
}

type ModifyHotfixesRequestXLanguageEnum struct {
	ZH_CN ModifyHotfixesRequestXLanguage
	EN_US ModifyHotfixesRequestXLanguage
}

func GetModifyHotfixesRequestXLanguageEnum() ModifyHotfixesRequestXLanguageEnum {
	return ModifyHotfixesRequestXLanguageEnum{
		ZH_CN: ModifyHotfixesRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ModifyHotfixesRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ModifyHotfixesRequestXLanguage) Value() string {
	return c.value
}

func (c ModifyHotfixesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifyHotfixesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
