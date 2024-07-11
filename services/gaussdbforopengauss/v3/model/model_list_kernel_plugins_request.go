package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListKernelPluginsRequest Request Object
type ListKernelPluginsRequest struct {

	// 语言。
	XLanguage *ListKernelPluginsRequestXLanguage `json:"X-Language,omitempty"`

	// 查询实例已安装的插件列表的实例ID
	InstanceId string `json:"instance_id"`
}

func (o ListKernelPluginsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKernelPluginsRequest struct{}"
	}

	return strings.Join([]string{"ListKernelPluginsRequest", string(data)}, " ")
}

type ListKernelPluginsRequestXLanguage struct {
	value string
}

type ListKernelPluginsRequestXLanguageEnum struct {
	ZH_CN ListKernelPluginsRequestXLanguage
	EN_US ListKernelPluginsRequestXLanguage
}

func GetListKernelPluginsRequestXLanguageEnum() ListKernelPluginsRequestXLanguageEnum {
	return ListKernelPluginsRequestXLanguageEnum{
		ZH_CN: ListKernelPluginsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListKernelPluginsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListKernelPluginsRequestXLanguage) Value() string {
	return c.value
}

func (c ListKernelPluginsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListKernelPluginsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
