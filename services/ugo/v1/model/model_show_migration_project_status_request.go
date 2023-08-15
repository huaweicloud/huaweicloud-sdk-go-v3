package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMigrationProjectStatusRequest Request Object
type ShowMigrationProjectStatusRequest struct {

	// 迁移项目ID。
	MigrationProjectId string `json:"migration_project_id"`
}

func (o ShowMigrationProjectStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMigrationProjectStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowMigrationProjectStatusRequest", string(data)}, " ")
}
