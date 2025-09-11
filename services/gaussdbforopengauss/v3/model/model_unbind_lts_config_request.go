package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UnbindLtsConfigRequest Request Object
type UnbindLtsConfigRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**:   - zh-cn   - en-us  **默认取值**: en-us
	XLanguage *UnbindLtsConfigRequestXLanguage `json:"X-Language,omitempty"`

	Body *UnbindLtsConfigRequestBody `json:"body,omitempty"`
}

func (o UnbindLtsConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnbindLtsConfigRequest struct{}"
	}

	return strings.Join([]string{"UnbindLtsConfigRequest", string(data)}, " ")
}

type UnbindLtsConfigRequestXLanguage struct {
	value string
}

type UnbindLtsConfigRequestXLanguageEnum struct {
	ZH_CN UnbindLtsConfigRequestXLanguage
	EN_US UnbindLtsConfigRequestXLanguage
}

func GetUnbindLtsConfigRequestXLanguageEnum() UnbindLtsConfigRequestXLanguageEnum {
	return UnbindLtsConfigRequestXLanguageEnum{
		ZH_CN: UnbindLtsConfigRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: UnbindLtsConfigRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c UnbindLtsConfigRequestXLanguage) Value() string {
	return c.value
}

func (c UnbindLtsConfigRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UnbindLtsConfigRequestXLanguage) UnmarshalJSON(b []byte) error {
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
