package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ChangeDeploymentSolutionRequest Request Object
type ChangeDeploymentSolutionRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn - en-us  **默认取值**: en-us
	XLanguage *ChangeDeploymentSolutionRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	Body *ChangeDeploymentSolutionRequestBody `json:"body,omitempty"`
}

func (o ChangeDeploymentSolutionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeDeploymentSolutionRequest struct{}"
	}

	return strings.Join([]string{"ChangeDeploymentSolutionRequest", string(data)}, " ")
}

type ChangeDeploymentSolutionRequestXLanguage struct {
	value string
}

type ChangeDeploymentSolutionRequestXLanguageEnum struct {
	ZH_CN ChangeDeploymentSolutionRequestXLanguage
	EN_US ChangeDeploymentSolutionRequestXLanguage
}

func GetChangeDeploymentSolutionRequestXLanguageEnum() ChangeDeploymentSolutionRequestXLanguageEnum {
	return ChangeDeploymentSolutionRequestXLanguageEnum{
		ZH_CN: ChangeDeploymentSolutionRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ChangeDeploymentSolutionRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ChangeDeploymentSolutionRequestXLanguage) Value() string {
	return c.value
}

func (c ChangeDeploymentSolutionRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChangeDeploymentSolutionRequestXLanguage) UnmarshalJSON(b []byte) error {
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
