package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPermissionCheckResultRequest Request Object
type ListPermissionCheckResultRequest struct {

	// 迁移项目ID。
	MigrationProjectId string `json:"migration_project_id"`

	// 分页查询的偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListPermissionCheckResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPermissionCheckResultRequest struct{}"
	}

	return strings.Join([]string{"ListPermissionCheckResultRequest", string(data)}, " ")
}
