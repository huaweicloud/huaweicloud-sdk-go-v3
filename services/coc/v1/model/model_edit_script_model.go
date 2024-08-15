package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EditScriptModel 修改脚本
type EditScriptModel struct {

	// 脚本描述
	Description string `json:"description"`

	// 脚本内容
	Content string `json:"content"`

	Properties *ScriptPropertiesModel `json:"properties,omitempty"`

	// 脚本入参
	ScriptParams *[]ScriptParamDefine `json:"script_params,omitempty"`
}

func (o EditScriptModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EditScriptModel struct{}"
	}

	return strings.Join([]string{"EditScriptModel", string(data)}, " ")
}
