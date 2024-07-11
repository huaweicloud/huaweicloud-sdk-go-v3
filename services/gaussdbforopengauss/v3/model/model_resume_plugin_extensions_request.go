package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ResumePluginExtensionsRequest Request Object
type ResumePluginExtensionsRequest struct {

	// 语言。
	XLanguage *ResumePluginExtensionsRequestXLanguage `json:"X-Language,omitempty"`

	// 配置插件拓展能力的实例ID
	InstanceId string `json:"instance_id"`

	Body *ResumePluginExtensionsRequestBody `json:"body,omitempty"`
}

func (o ResumePluginExtensionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResumePluginExtensionsRequest struct{}"
	}

	return strings.Join([]string{"ResumePluginExtensionsRequest", string(data)}, " ")
}

type ResumePluginExtensionsRequestXLanguage struct {
	value string
}

type ResumePluginExtensionsRequestXLanguageEnum struct {
	ZH_CN ResumePluginExtensionsRequestXLanguage
	EN_US ResumePluginExtensionsRequestXLanguage
}

func GetResumePluginExtensionsRequestXLanguageEnum() ResumePluginExtensionsRequestXLanguageEnum {
	return ResumePluginExtensionsRequestXLanguageEnum{
		ZH_CN: ResumePluginExtensionsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ResumePluginExtensionsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ResumePluginExtensionsRequestXLanguage) Value() string {
	return c.value
}

func (c ResumePluginExtensionsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResumePluginExtensionsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
