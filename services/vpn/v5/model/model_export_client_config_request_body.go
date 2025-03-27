package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ExportClientConfigRequestBody struct {

	// 操作系统类型
	OsType *ExportClientConfigRequestBodyOsType `json:"os_type,omitempty"`
}

func (o ExportClientConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportClientConfigRequestBody struct{}"
	}

	return strings.Join([]string{"ExportClientConfigRequestBody", string(data)}, " ")
}

type ExportClientConfigRequestBodyOsType struct {
	value string
}

type ExportClientConfigRequestBodyOsTypeEnum struct {
	WINDOWS ExportClientConfigRequestBodyOsType
	LINUX   ExportClientConfigRequestBodyOsType
	MAC_OS  ExportClientConfigRequestBodyOsType
	ANDROID ExportClientConfigRequestBodyOsType
	I_OS    ExportClientConfigRequestBodyOsType
}

func GetExportClientConfigRequestBodyOsTypeEnum() ExportClientConfigRequestBodyOsTypeEnum {
	return ExportClientConfigRequestBodyOsTypeEnum{
		WINDOWS: ExportClientConfigRequestBodyOsType{
			value: "Windows",
		},
		LINUX: ExportClientConfigRequestBodyOsType{
			value: "Linux",
		},
		MAC_OS: ExportClientConfigRequestBodyOsType{
			value: "MacOS",
		},
		ANDROID: ExportClientConfigRequestBodyOsType{
			value: "Android",
		},
		I_OS: ExportClientConfigRequestBodyOsType{
			value: "iOS",
		},
	}
}

func (c ExportClientConfigRequestBodyOsType) Value() string {
	return c.value
}

func (c ExportClientConfigRequestBodyOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportClientConfigRequestBodyOsType) UnmarshalJSON(b []byte) error {
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
