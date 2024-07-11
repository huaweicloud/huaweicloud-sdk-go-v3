package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListPluginExtensionsRequest Request Object
type ListPluginExtensionsRequest struct {

	// 语言。
	XLanguage *ListPluginExtensionsRequestXLanguage `json:"X-Language,omitempty"`

	// 查询实例插件拓展信息的实例ID
	InstanceId string `json:"instance_id"`

	Body *ListPluginExtensionsRequestBody `json:"body,omitempty"`
}

func (o ListPluginExtensionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPluginExtensionsRequest struct{}"
	}

	return strings.Join([]string{"ListPluginExtensionsRequest", string(data)}, " ")
}

type ListPluginExtensionsRequestXLanguage struct {
	value string
}

type ListPluginExtensionsRequestXLanguageEnum struct {
	ZH_CN ListPluginExtensionsRequestXLanguage
	EN_US ListPluginExtensionsRequestXLanguage
}

func GetListPluginExtensionsRequestXLanguageEnum() ListPluginExtensionsRequestXLanguageEnum {
	return ListPluginExtensionsRequestXLanguageEnum{
		ZH_CN: ListPluginExtensionsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListPluginExtensionsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListPluginExtensionsRequestXLanguage) Value() string {
	return c.value
}

func (c ListPluginExtensionsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPluginExtensionsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
