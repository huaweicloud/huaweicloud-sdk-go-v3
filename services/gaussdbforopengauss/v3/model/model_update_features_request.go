package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateFeaturesRequest Request Object
type UpdateFeaturesRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *UpdateFeaturesRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	Body *UpdateFeaturesRequestBody `json:"body,omitempty"`
}

func (o UpdateFeaturesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFeaturesRequest struct{}"
	}

	return strings.Join([]string{"UpdateFeaturesRequest", string(data)}, " ")
}

type UpdateFeaturesRequestXLanguage struct {
	value string
}

type UpdateFeaturesRequestXLanguageEnum struct {
	ZH_CN UpdateFeaturesRequestXLanguage
	EN_US UpdateFeaturesRequestXLanguage
}

func GetUpdateFeaturesRequestXLanguageEnum() UpdateFeaturesRequestXLanguageEnum {
	return UpdateFeaturesRequestXLanguageEnum{
		ZH_CN: UpdateFeaturesRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: UpdateFeaturesRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c UpdateFeaturesRequestXLanguage) Value() string {
	return c.value
}

func (c UpdateFeaturesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateFeaturesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
