package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMigrationProject 创建迁移项目请求体。
type CreateMigrationProject struct {

	// 迁移项目名称。长度为5-50个字符，以英文字母开头，英文字母或数字结束，允许包含下划线和中划线。不允许重复。
	MigrationProjectName string `json:"migration_project_name"`

	// 评估项目ID。
	EvaluationProjectId int32 `json:"evaluation_project_id"`

	TargetDbInfo *TargetDbInfo `json:"target_db_info"`

	OpenGaussConfig *OpenGaussConfig `json:"open_gauss_config,omitempty"`
}

func (o CreateMigrationProject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMigrationProject struct{}"
	}

	return strings.Join([]string{"CreateMigrationProject", string(data)}, " ")
}
