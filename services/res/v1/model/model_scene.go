package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Scene
type Scene struct {

	// 类型。
	Category string `json:"category"`

	// 场景类型。
	Type string `json:"type"`

	// 场景名称。
	SceneName string `json:"scene_name"`

	// 场景id。
	SceneId string `json:"scene_id"`

	// 数据源id。
	DatasourceId string `json:"datasource_id"`

	// 状态。
	Status string `json:"status"`

	// 创建时间。
	CreatedAt int64 `json:"created_at"`

	// 更新时间。
	UpdateAt *int64 `json:"update_at,omitempty"`

	// 工作空间id。
	WorkspaceId string `json:"workspace_id"`

	// 服务类型。
	ServiceType string `json:"service_type"`
}

func (o Scene) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Scene struct{}"
	}

	return strings.Join([]string{"Scene", string(data)}, " ")
}
