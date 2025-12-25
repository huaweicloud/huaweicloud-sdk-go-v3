package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AnalysisScript struct {

	// UUID
	ScriptId *string `json:"script_id,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 工作空间ID
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 脚本名称
	ScriptName *string `json:"script_name,omitempty"`

	Category *ScriptCategory `json:"category,omitempty"`

	// 脚本目录分组名称，长度在1到256个字符之间。
	Directory *string `json:"directory,omitempty"`

	// 脚本的相关描述信息，长度在1到1024个字符之间。
	Description *string `json:"description,omitempty"`

	ScriptType *AnalysisScriptType `json:"script_type,omitempty"`

	// UUID
	RetrieveTableId *string `json:"retrieve_table_id,omitempty"`

	// 脚本内容，长度在1到10240个字符之间。
	Script *string `json:"script,omitempty"`

	// Iam用户ID
	Owner *string `json:"owner,omitempty"`

	// 分析脚本参数列表
	ScriptParams *[]AnalysisScriptParam `json:"script_params,omitempty"`

	// Iam用户ID
	CreateBy *string `json:"create_by,omitempty"`

	// 毫秒时间戳
	CreateTime *int64 `json:"create_time,omitempty"`

	// Iam用户ID
	UpdateBy *string `json:"update_by,omitempty"`

	// 毫秒时间戳
	UpdateTime *int64 `json:"update_time,omitempty"`
}

func (o AnalysisScript) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AnalysisScript struct{}"
	}

	return strings.Join([]string{"AnalysisScript", string(data)}, " ")
}
