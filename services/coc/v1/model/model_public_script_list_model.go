package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PublicScriptListModel 公共脚本列表对象
type PublicScriptListModel struct {

	// 脚本自增id
	Id int64 `json:"id"`

	// 脚本uuid
	ScriptUuid string `json:"script_uuid"`

	// 脚本名称
	Name string `json:"name"`

	// 脚本类型 SHELL:shell脚本 PYTHON:python脚本 BAT:bat脚本
	Type PublicScriptListModelType `json:"type"`

	// 创建时间
	GmtCreated int64 `json:"gmt_created"`

	// 脚本描述： 根据X-Language(X-Language) 进行国际化
	Description string `json:"description"`

	Properties *PublicScriptPropertiesModel `json:"properties"`
}

func (o PublicScriptListModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicScriptListModel struct{}"
	}

	return strings.Join([]string{"PublicScriptListModel", string(data)}, " ")
}

type PublicScriptListModelType struct {
	value string
}

type PublicScriptListModelTypeEnum struct {
	SHELL  PublicScriptListModelType
	PYTHON PublicScriptListModelType
	BAT    PublicScriptListModelType
}

func GetPublicScriptListModelTypeEnum() PublicScriptListModelTypeEnum {
	return PublicScriptListModelTypeEnum{
		SHELL: PublicScriptListModelType{
			value: "SHELL",
		},
		PYTHON: PublicScriptListModelType{
			value: "PYTHON",
		},
		BAT: PublicScriptListModelType{
			value: "BAT",
		},
	}
}

func (c PublicScriptListModelType) Value() string {
	return c.value
}

func (c PublicScriptListModelType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PublicScriptListModelType) UnmarshalJSON(b []byte) error {
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
