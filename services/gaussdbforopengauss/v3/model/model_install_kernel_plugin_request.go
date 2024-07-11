package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// InstallKernelPluginRequest Request Object
type InstallKernelPluginRequest struct {

	// 语言。
	XLanguage *InstallKernelPluginRequestXLanguage `json:"X-Language,omitempty"`

	// 需要安装插件的实例id
	InstanceId string `json:"instance_id"`

	Body *InstallKernelPluginRequestBody `json:"body,omitempty"`
}

func (o InstallKernelPluginRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstallKernelPluginRequest struct{}"
	}

	return strings.Join([]string{"InstallKernelPluginRequest", string(data)}, " ")
}

type InstallKernelPluginRequestXLanguage struct {
	value string
}

type InstallKernelPluginRequestXLanguageEnum struct {
	ZH_CN InstallKernelPluginRequestXLanguage
	EN_US InstallKernelPluginRequestXLanguage
}

func GetInstallKernelPluginRequestXLanguageEnum() InstallKernelPluginRequestXLanguageEnum {
	return InstallKernelPluginRequestXLanguageEnum{
		ZH_CN: InstallKernelPluginRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: InstallKernelPluginRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c InstallKernelPluginRequestXLanguage) Value() string {
	return c.value
}

func (c InstallKernelPluginRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstallKernelPluginRequestXLanguage) UnmarshalJSON(b []byte) error {
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
