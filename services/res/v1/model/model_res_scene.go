package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResScene struct {

	// 类型。
	Category *string `json:"category,omitempty"`

	// 数据源id。
	DatasourceId *string `json:"datasource_id,omitempty"`

	DsConfig *DataConfig `json:"ds_config,omitempty"`

	// 场景id。
	SceneId *string `json:"scene_id,omitempty"`

	// 场景名称。
	SceneName *string `json:"scene_name,omitempty"`

	// 场景类型。
	Type *string `json:"type,omitempty"`

	// 服务类型。
	ServiceType *string `json:"service_type,omitempty"`

	// 状态。
	Status *string `json:"status,omitempty"`

	// 工作空间id。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 创建时间。
	CreatedAt *int64 `json:"created_at,omitempty"`

	// 更新时间。
	UpdateAt *int64 `json:"update_at,omitempty"`

	SpecsConfig *SpecsConfig `json:"specs_config,omitempty"`
}

func (o ResScene) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResScene struct{}"
	}

	return strings.Join([]string{"ResScene", string(data)}, " ")
}
