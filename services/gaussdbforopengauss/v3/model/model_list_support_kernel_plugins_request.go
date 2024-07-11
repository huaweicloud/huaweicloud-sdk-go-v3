package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSupportKernelPluginsRequest Request Object
type ListSupportKernelPluginsRequest struct {

	// 语言。
	XLanguage *ListSupportKernelPluginsRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ListSupportKernelPluginsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSupportKernelPluginsRequest struct{}"
	}

	return strings.Join([]string{"ListSupportKernelPluginsRequest", string(data)}, " ")
}

type ListSupportKernelPluginsRequestXLanguage struct {
	value string
}

type ListSupportKernelPluginsRequestXLanguageEnum struct {
	ZH_CN ListSupportKernelPluginsRequestXLanguage
	EN_US ListSupportKernelPluginsRequestXLanguage
}

func GetListSupportKernelPluginsRequestXLanguageEnum() ListSupportKernelPluginsRequestXLanguageEnum {
	return ListSupportKernelPluginsRequestXLanguageEnum{
		ZH_CN: ListSupportKernelPluginsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListSupportKernelPluginsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListSupportKernelPluginsRequestXLanguage) Value() string {
	return c.value
}

func (c ListSupportKernelPluginsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSupportKernelPluginsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
