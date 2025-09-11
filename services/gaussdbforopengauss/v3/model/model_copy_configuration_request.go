package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CopyConfigurationRequest Request Object
type CopyConfigurationRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**:   - zh-cn   - en-us  **默认取值**: en-us
	XLanguage *CopyConfigurationRequestXLanguage `json:"X-Language,omitempty"`

	// 被复制的参数模板ID。
	ConfigId string `json:"config_id"`

	Body *ParamGroupCopyRequestBody `json:"body,omitempty"`
}

func (o CopyConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyConfigurationRequest struct{}"
	}

	return strings.Join([]string{"CopyConfigurationRequest", string(data)}, " ")
}

type CopyConfigurationRequestXLanguage struct {
	value string
}

type CopyConfigurationRequestXLanguageEnum struct {
	ZH_CN CopyConfigurationRequestXLanguage
	EN_US CopyConfigurationRequestXLanguage
}

func GetCopyConfigurationRequestXLanguageEnum() CopyConfigurationRequestXLanguageEnum {
	return CopyConfigurationRequestXLanguageEnum{
		ZH_CN: CopyConfigurationRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: CopyConfigurationRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c CopyConfigurationRequestXLanguage) Value() string {
	return c.value
}

func (c CopyConfigurationRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CopyConfigurationRequestXLanguage) UnmarshalJSON(b []byte) error {
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
