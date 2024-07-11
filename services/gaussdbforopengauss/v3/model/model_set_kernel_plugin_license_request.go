package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SetKernelPluginLicenseRequest Request Object
type SetKernelPluginLicenseRequest struct {

	// 语言。
	XLanguage *SetKernelPluginLicenseRequestXLanguage `json:"X-Language,omitempty"`

	// 需要配置license的实例
	InstanceId string `json:"instance_id"`

	Body *SetKernelPluginLicenseRequestBody `json:"body,omitempty"`
}

func (o SetKernelPluginLicenseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetKernelPluginLicenseRequest struct{}"
	}

	return strings.Join([]string{"SetKernelPluginLicenseRequest", string(data)}, " ")
}

type SetKernelPluginLicenseRequestXLanguage struct {
	value string
}

type SetKernelPluginLicenseRequestXLanguageEnum struct {
	ZH_CN SetKernelPluginLicenseRequestXLanguage
	EN_US SetKernelPluginLicenseRequestXLanguage
}

func GetSetKernelPluginLicenseRequestXLanguageEnum() SetKernelPluginLicenseRequestXLanguageEnum {
	return SetKernelPluginLicenseRequestXLanguageEnum{
		ZH_CN: SetKernelPluginLicenseRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: SetKernelPluginLicenseRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c SetKernelPluginLicenseRequestXLanguage) Value() string {
	return c.value
}

func (c SetKernelPluginLicenseRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SetKernelPluginLicenseRequestXLanguage) UnmarshalJSON(b []byte) error {
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
