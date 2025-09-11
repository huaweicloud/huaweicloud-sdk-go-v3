package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BindLtsConfigRequest Request Object
type BindLtsConfigRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**:   - zh-cn   - en-us  **默认取值**: en-us
	XLanguage *BindLtsConfigRequestXLanguage `json:"X-Language,omitempty"`

	Body *BindLtsConfigRequestBody `json:"body,omitempty"`
}

func (o BindLtsConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindLtsConfigRequest struct{}"
	}

	return strings.Join([]string{"BindLtsConfigRequest", string(data)}, " ")
}

type BindLtsConfigRequestXLanguage struct {
	value string
}

type BindLtsConfigRequestXLanguageEnum struct {
	ZH_CN BindLtsConfigRequestXLanguage
	EN_US BindLtsConfigRequestXLanguage
}

func GetBindLtsConfigRequestXLanguageEnum() BindLtsConfigRequestXLanguageEnum {
	return BindLtsConfigRequestXLanguageEnum{
		ZH_CN: BindLtsConfigRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: BindLtsConfigRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c BindLtsConfigRequestXLanguage) Value() string {
	return c.value
}

func (c BindLtsConfigRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BindLtsConfigRequestXLanguage) UnmarshalJSON(b []byte) error {
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
