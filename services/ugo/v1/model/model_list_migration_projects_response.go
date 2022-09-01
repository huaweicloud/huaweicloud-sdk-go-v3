package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListMigrationProjectsResponse struct {

	// 当前页的迁移项目列表。
	MigrationProjects *[]MigrationProject `json:"migration_projects,omitempty" xml:"migration_projects"`

	// 迁移项目总数。
	TotalCount     *int32 `json:"total_count,omitempty" xml:"total_count"`
	HttpStatusCode int    `json:"-"`
}

func (o ListMigrationProjectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMigrationProjectsResponse struct{}"
	}

	return strings.Join([]string{"ListMigrationProjectsResponse", string(data)}, " ")
}
