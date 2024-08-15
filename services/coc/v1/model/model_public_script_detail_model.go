package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PublicScriptDetailModel 公共脚本详情
type PublicScriptDetailModel struct {

	// 脚本uuid
	ScriptUuid string `json:"script_uuid"`

	// 脚本名称
	Name string `json:"name"`

	// 脚本描述
	Description string `json:"description"`

	// 脚本类型 SHELL:shell脚本， PYTHON:Python脚本， BAT:Bat脚本，
	Type PublicScriptDetailModelType `json:"type"`

	// 脚本内容
	Content string `json:"content"`

	// 脚本入参
	ScriptParams *[]ScriptParamDefine `json:"script_params,omitempty"`

	// 创建时间
	GmtCreated int64 `json:"gmt_created"`

	Properties *PublicScriptPropertiesModel `json:"properties"`
}

func (o PublicScriptDetailModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicScriptDetailModel struct{}"
	}

	return strings.Join([]string{"PublicScriptDetailModel", string(data)}, " ")
}

type PublicScriptDetailModelType struct {
	value string
}

type PublicScriptDetailModelTypeEnum struct {
	SHELL  PublicScriptDetailModelType
	PYTHON PublicScriptDetailModelType
	BAT    PublicScriptDetailModelType
}

func GetPublicScriptDetailModelTypeEnum() PublicScriptDetailModelTypeEnum {
	return PublicScriptDetailModelTypeEnum{
		SHELL: PublicScriptDetailModelType{
			value: "SHELL",
		},
		PYTHON: PublicScriptDetailModelType{
			value: "PYTHON",
		},
		BAT: PublicScriptDetailModelType{
			value: "BAT",
		},
	}
}

func (c PublicScriptDetailModelType) Value() string {
	return c.value
}

func (c PublicScriptDetailModelType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PublicScriptDetailModelType) UnmarshalJSON(b []byte) error {
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
