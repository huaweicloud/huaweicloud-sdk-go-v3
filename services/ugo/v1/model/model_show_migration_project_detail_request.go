package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowMigrationProjectDetailRequest struct {

	// 迁移项目ID。
	MigrationProjectId string `json:"migration_project_id" xml:"migration_project_id"`
}

func (o ShowMigrationProjectDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMigrationProjectDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowMigrationProjectDetailRequest", string(data)}, " ")
}
