package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 迁移资源
type MigrateResource struct {
	// 项目ID。resource_type为region级别服务时为必选项。

	ProjectId *string `json:"project_id,omitempty"`
	// 资源ID

	ResourceId string `json:"resource_id"`
	// 资源类型

	ResourceType string `json:"resource_type"`
	// 是否关联迁移。目前仅支持ECS关联资源EVS、EIP迁移。

	Associated *bool `json:"associated,omitempty"`
}

func (o MigrateResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateResource struct{}"
	}

	return strings.Join([]string{"MigrateResource", string(data)}, " ")
}
