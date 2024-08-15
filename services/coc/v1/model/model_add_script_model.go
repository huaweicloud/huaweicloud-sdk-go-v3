package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AddScriptModel **约束：** - 1. 同poject_id下不允许同名脚本 - 2. 修改脚本时，不允许修改：脚本名称、project_id - 3. 字段长度和范围限制  - 脚本内容长度，创建时为做校验
type AddScriptModel struct {

	// 脚本名称：只能包含中文、英文、数字、下划线
	Name string `json:"name"`

	Properties *ScriptPropertiesModel `json:"properties"`

	// 脚本描述
	Description string `json:"description"`

	// 脚本类型: 对于脚本后缀： SHELL:.sh PYTHON:.py BAT:.bat
	Type AddScriptModelType `json:"type"`

	// 脚本内容
	Content string `json:"content"`

	// 企业项目ID，默认为：0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 脚本入参
	ScriptParams *[]ScriptParamDefine `json:"script_params,omitempty"`
}

func (o AddScriptModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddScriptModel struct{}"
	}

	return strings.Join([]string{"AddScriptModel", string(data)}, " ")
}

type AddScriptModelType struct {
	value string
}

type AddScriptModelTypeEnum struct {
	SHELL  AddScriptModelType
	PYTHON AddScriptModelType
	BAT    AddScriptModelType
}

func GetAddScriptModelTypeEnum() AddScriptModelTypeEnum {
	return AddScriptModelTypeEnum{
		SHELL: AddScriptModelType{
			value: "SHELL",
		},
		PYTHON: AddScriptModelType{
			value: "PYTHON",
		},
		BAT: AddScriptModelType{
			value: "BAT",
		},
	}
}

func (c AddScriptModelType) Value() string {
	return c.value
}

func (c AddScriptModelType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddScriptModelType) UnmarshalJSON(b []byte) error {
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
