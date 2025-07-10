package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateScriptReq 创建脚本信息。
type CreateScriptReq struct {

	// 脚本名称。
	Name string `json:"name"`

	// 脚本类型：POWERSHELL/BAT/SHELL。
	Type CreateScriptReqType `json:"type"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 脚本内容。
	Content string `json:"content"`

	// 脚本版本。
	Version *string `json:"version,omitempty"`
}

func (o CreateScriptReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScriptReq struct{}"
	}

	return strings.Join([]string{"CreateScriptReq", string(data)}, " ")
}

type CreateScriptReqType struct {
	value string
}

type CreateScriptReqTypeEnum struct {
	POWERSHELL CreateScriptReqType
	BAT        CreateScriptReqType
	SHELL      CreateScriptReqType
}

func GetCreateScriptReqTypeEnum() CreateScriptReqTypeEnum {
	return CreateScriptReqTypeEnum{
		POWERSHELL: CreateScriptReqType{
			value: "POWERSHELL",
		},
		BAT: CreateScriptReqType{
			value: "BAT",
		},
		SHELL: CreateScriptReqType{
			value: "SHELL",
		},
	}
}

func (c CreateScriptReqType) Value() string {
	return c.value
}

func (c CreateScriptReqType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateScriptReqType) UnmarshalJSON(b []byte) error {
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
