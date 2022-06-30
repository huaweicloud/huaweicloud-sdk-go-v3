package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 单个评估项目。
type EvaluationProject struct {

	// 评估项目ID。
	EvaluationProjectId int32 `json:"evaluation_project_id"`

	// 评估项目名称。
	EvaluationProjectName string `json:"evaluation_project_name"`

	// 评估项目状态。
	EvaluationProjectStatus string `json:"evaluation_project_status"`

	ProjectStatusDetail *ProjectStatusDetail `json:"project_status_detail,omitempty"`

	// 源数据库类型。
	SourceDbType string `json:"source_db_type"`

	// 源数据库版本。
	SourceDbVersion string `json:"source_db_version"`

	// 目标数据库类型。
	TargetDbType *string `json:"target_db_type,omitempty"`

	// 目标数据库版本。
	TargetDbVersion *string `json:"target_db_version,omitempty"`

	// 已收集的SQL大小，单位：B。
	CollectSize *int64 `json:"collect_size,omitempty"`

	// 资源ID。
	ResourceId string `json:"resource_id"`

	// 创建时间。
	CreatedTime string `json:"created_time"`

	// 更新时间。
	UpdatedTime string `json:"updated_time"`

	// 失败原因。
	ErrorReason *string `json:"error_reason,omitempty"`
}

func (o EvaluationProject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EvaluationProject struct{}"
	}

	return strings.Join([]string{"EvaluationProject", string(data)}, " ")
}
