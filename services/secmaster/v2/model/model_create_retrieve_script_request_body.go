package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateRetrieveScriptRequestBody struct {

	// UUID
	TableId string `json:"table_id"`

	// 脚本名称
	ScriptName string `json:"script_name"`

	Category *ScriptCategory `json:"category"`

	// 脚本目录分组名称，长度在1到256个字符之间。
	Directory *string `json:"directory,omitempty"`

	// 脚本的相关描述信息，长度在1到1024个字符之间。
	Description *string `json:"description,omitempty"`

	// 脚本内容，长度在1到10240个字符之间。
	Script string `json:"script"`
}

func (o CreateRetrieveScriptRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRetrieveScriptRequestBody struct{}"
	}

	return strings.Join([]string{"CreateRetrieveScriptRequestBody", string(data)}, " ")
}
