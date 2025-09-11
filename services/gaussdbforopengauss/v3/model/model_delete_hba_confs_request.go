package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteHbaConfsRequest Request Object
type DeleteHbaConfsRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *DeleteHbaConfsRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	Body *DeleteHbaConfRequestBody `json:"body,omitempty"`
}

func (o DeleteHbaConfsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHbaConfsRequest struct{}"
	}

	return strings.Join([]string{"DeleteHbaConfsRequest", string(data)}, " ")
}

type DeleteHbaConfsRequestXLanguage struct {
	value string
}

type DeleteHbaConfsRequestXLanguageEnum struct {
	ZH_CN DeleteHbaConfsRequestXLanguage
	EN_US DeleteHbaConfsRequestXLanguage
}

func GetDeleteHbaConfsRequestXLanguageEnum() DeleteHbaConfsRequestXLanguageEnum {
	return DeleteHbaConfsRequestXLanguageEnum{
		ZH_CN: DeleteHbaConfsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: DeleteHbaConfsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c DeleteHbaConfsRequestXLanguage) Value() string {
	return c.value
}

func (c DeleteHbaConfsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteHbaConfsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
