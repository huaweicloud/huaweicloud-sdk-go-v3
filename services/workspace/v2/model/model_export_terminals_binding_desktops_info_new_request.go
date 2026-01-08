package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExportTerminalsBindingDesktopsInfoNewRequest Request Object
type ExportTerminalsBindingDesktopsInfoNewRequest struct {

	// 语言。  - zh_CN：中文 - en_US：英文
	Language ExportTerminalsBindingDesktopsInfoNewRequestLanguage `json:"language"`

	// 桌面名。
	ComputerName *string `json:"computer_name,omitempty"`

	// mac地址。
	Mac *string `json:"mac,omitempty"`
}

func (o ExportTerminalsBindingDesktopsInfoNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportTerminalsBindingDesktopsInfoNewRequest struct{}"
	}

	return strings.Join([]string{"ExportTerminalsBindingDesktopsInfoNewRequest", string(data)}, " ")
}

type ExportTerminalsBindingDesktopsInfoNewRequestLanguage struct {
	value string
}

type ExportTerminalsBindingDesktopsInfoNewRequestLanguageEnum struct {
	ZH_CN ExportTerminalsBindingDesktopsInfoNewRequestLanguage
	EN_US ExportTerminalsBindingDesktopsInfoNewRequestLanguage
}

func GetExportTerminalsBindingDesktopsInfoNewRequestLanguageEnum() ExportTerminalsBindingDesktopsInfoNewRequestLanguageEnum {
	return ExportTerminalsBindingDesktopsInfoNewRequestLanguageEnum{
		ZH_CN: ExportTerminalsBindingDesktopsInfoNewRequestLanguage{
			value: "zh_CN",
		},
		EN_US: ExportTerminalsBindingDesktopsInfoNewRequestLanguage{
			value: "en_US",
		},
	}
}

func (c ExportTerminalsBindingDesktopsInfoNewRequestLanguage) Value() string {
	return c.value
}

func (c ExportTerminalsBindingDesktopsInfoNewRequestLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportTerminalsBindingDesktopsInfoNewRequestLanguage) UnmarshalJSON(b []byte) error {
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
