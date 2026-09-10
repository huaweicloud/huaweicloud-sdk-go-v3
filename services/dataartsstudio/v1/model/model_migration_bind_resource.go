package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MigrationBindResource 待关联或取消关联的数据集成资源信息。
type MigrationBindResource struct {

	// 资源ID。
	ResourceId string `json:"resource_id"`

	// 资源名称。
	ResourceName string `json:"resource_name"`

	// 资源关联的工作空间列表。
	Workspaces []MigrationBindWorkspace `json:"workspaces"`
}

func (o MigrationBindResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrationBindResource struct{}"
	}

	return strings.Join([]string{"MigrationBindResource", string(data)}, " ")
}
