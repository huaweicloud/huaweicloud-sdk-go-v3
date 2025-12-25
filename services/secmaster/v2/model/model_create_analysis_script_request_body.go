package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateAnalysisScriptRequestBody struct {

	// 脚本名称
	ScriptName string `json:"script_name"`

	Category *ScriptCategory `json:"category"`

	// 脚本目录分组名称，长度在1到256个字符之间。
	Directory *string `json:"directory,omitempty"`

	// 脚本的相关描述信息，长度在1到1024个字符之间。
	Description *string `json:"description,omitempty"`

	ScriptType *AnalysisScriptType `json:"script_type"`

	// UUID
	RetrieveTableId *string `json:"retrieve_table_id,omitempty"`

	// 表名
	RetrieveTableName *string `json:"retrieve_table_name,omitempty"`

	// 脚本内容，长度在1到10240个字符之间。
	Script string `json:"script"`

	// Iam用户ID
	Owner *string `json:"owner,omitempty"`

	// 分析脚本参数列表
	ScriptParams []AnalysisScriptParam `json:"script_params"`
}

func (o CreateAnalysisScriptRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAnalysisScriptRequestBody struct{}"
	}

	return strings.Join([]string{"CreateAnalysisScriptRequestBody", string(data)}, " ")
}
