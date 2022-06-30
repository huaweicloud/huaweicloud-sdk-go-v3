package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowMigrationProjectDetailResponse struct {

	// 迁移项目ID。
	MigrationProjectId *int32 `json:"migration_project_id,omitempty"`

	// 迁移项目状态。
	MigrationProjectName *string `json:"migration_project_name,omitempty"`

	// 对应的评估项目名称。
	EvaluationProjectName *string `json:"evaluation_project_name,omitempty"`

	SourceDbInfo *DataBase `json:"source_db_info,omitempty"`

	TargetDbInfo *DataBase `json:"target_db_info,omitempty"`

	// 创建时间。
	CreatedTime *string `json:"created_time,omitempty"`

	// 更新时间。
	UpdatedTime    *string `json:"updated_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowMigrationProjectDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMigrationProjectDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowMigrationProjectDetailResponse", string(data)}, " ")
}
