package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteMigrationProjectResponse Response Object
type DeleteMigrationProjectResponse struct {

	// 迁移项目ID。
	MigrationProjectId *string `json:"migration_project_id,omitempty"`
	HttpStatusCode     int     `json:"-"`
}

func (o DeleteMigrationProjectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMigrationProjectResponse struct{}"
	}

	return strings.Join([]string{"DeleteMigrationProjectResponse", string(data)}, " ")
}
