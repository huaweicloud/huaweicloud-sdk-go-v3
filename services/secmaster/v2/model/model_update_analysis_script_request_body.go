package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateAnalysisScriptRequestBody struct {

	// 脚本名称
	ScriptName string `json:"script_name"`

	// 脚本目录分组名称，长度在1到256个字符之间。
	Directory *string `json:"directory,omitempty"`

	// 脚本的相关描述信息，长度在1到1024个字符之间。
	Description *string `json:"description,omitempty"`

	ScriptType *AnalysisScriptType `json:"script_type"`

	// UUID
	RetrieveTableId *string `json:"retrieve_table_id,omitempty"`

	// 脚本内容，长度在1到10240个字符之间。
	Script string `json:"script"`

	// 分析脚本参数列表
	Owner *[]AnalysisScriptParam `json:"owner,omitempty"`

	// 分析脚本参数列表
	ScriptParams []AnalysisScriptParam `json:"script_params"`
}

func (o UpdateAnalysisScriptRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAnalysisScriptRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAnalysisScriptRequestBody", string(data)}, " ")
}
