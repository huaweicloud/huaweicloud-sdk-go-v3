package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AddHbaConfsRequest Request Object
type AddHbaConfsRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *AddHbaConfsRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	Body *AddHbaConfRequestBody `json:"body,omitempty"`
}

func (o AddHbaConfsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddHbaConfsRequest struct{}"
	}

	return strings.Join([]string{"AddHbaConfsRequest", string(data)}, " ")
}

type AddHbaConfsRequestXLanguage struct {
	value string
}

type AddHbaConfsRequestXLanguageEnum struct {
	ZH_CN AddHbaConfsRequestXLanguage
	EN_US AddHbaConfsRequestXLanguage
}

func GetAddHbaConfsRequestXLanguageEnum() AddHbaConfsRequestXLanguageEnum {
	return AddHbaConfsRequestXLanguageEnum{
		ZH_CN: AddHbaConfsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: AddHbaConfsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c AddHbaConfsRequestXLanguage) Value() string {
	return c.value
}

func (c AddHbaConfsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddHbaConfsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
