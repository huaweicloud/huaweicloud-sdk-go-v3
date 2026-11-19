package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// InstallPluginRequest Request Object
type InstallPluginRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn - en-us  **默认取值**: en-us
	XLanguage *InstallPluginRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	Body *InstallPluginRequestBody `json:"body,omitempty"`
}

func (o InstallPluginRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstallPluginRequest struct{}"
	}

	return strings.Join([]string{"InstallPluginRequest", string(data)}, " ")
}

type InstallPluginRequestXLanguage struct {
	value string
}

type InstallPluginRequestXLanguageEnum struct {
	ZH_CN InstallPluginRequestXLanguage
	EN_US InstallPluginRequestXLanguage
}

func GetInstallPluginRequestXLanguageEnum() InstallPluginRequestXLanguageEnum {
	return InstallPluginRequestXLanguageEnum{
		ZH_CN: InstallPluginRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: InstallPluginRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c InstallPluginRequestXLanguage) Value() string {
	return c.value
}

func (c InstallPluginRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstallPluginRequestXLanguage) UnmarshalJSON(b []byte) error {
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
