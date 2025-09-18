package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateRedistributionControlRequest Request Object
type UpdateRedistributionControlRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *UpdateRedistributionControlRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及 **取值范围**: 只能由英文字母、数字组成，且长度为32个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	Body *RedistributionRequestBody `json:"body,omitempty"`
}

func (o UpdateRedistributionControlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRedistributionControlRequest struct{}"
	}

	return strings.Join([]string{"UpdateRedistributionControlRequest", string(data)}, " ")
}

type UpdateRedistributionControlRequestXLanguage struct {
	value string
}

type UpdateRedistributionControlRequestXLanguageEnum struct {
	ZH_CN UpdateRedistributionControlRequestXLanguage
	EN_US UpdateRedistributionControlRequestXLanguage
}

func GetUpdateRedistributionControlRequestXLanguageEnum() UpdateRedistributionControlRequestXLanguageEnum {
	return UpdateRedistributionControlRequestXLanguageEnum{
		ZH_CN: UpdateRedistributionControlRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: UpdateRedistributionControlRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c UpdateRedistributionControlRequestXLanguage) Value() string {
	return c.value
}

func (c UpdateRedistributionControlRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateRedistributionControlRequestXLanguage) UnmarshalJSON(b []byte) error {
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
