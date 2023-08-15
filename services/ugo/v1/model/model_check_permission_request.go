package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckPermissionRequest Request Object
type CheckPermissionRequest struct {

	// 迁移项目ID。
	MigrationProjectId string `json:"migration_project_id"`
}

func (o CheckPermissionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckPermissionRequest struct{}"
	}

	return strings.Join([]string{"CheckPermissionRequest", string(data)}, " ")
}
