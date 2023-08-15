package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteMigrationProjectRequest Request Object
type DeleteMigrationProjectRequest struct {

	// 迁移项目ID。
	MigrationProjectId string `json:"migration_project_id"`
}

func (o DeleteMigrationProjectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMigrationProjectRequest struct{}"
	}

	return strings.Join([]string{"DeleteMigrationProjectRequest", string(data)}, " ")
}
