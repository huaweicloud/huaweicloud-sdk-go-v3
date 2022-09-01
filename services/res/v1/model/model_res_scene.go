package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResScene struct {

	// 类型。
	Category *string `json:"category,omitempty" xml:"category"`

	// 数据源id。
	DatasourceId *string `json:"datasource_id,omitempty" xml:"datasource_id"`

	DsConfig *DataConfig `json:"ds_config,omitempty" xml:"ds_config"`

	// 场景id。
	SceneId *string `json:"scene_id,omitempty" xml:"scene_id"`

	// 场景名称。
	SceneName *string `json:"scene_name,omitempty" xml:"scene_name"`

	// 场景类型。
	Type *string `json:"type,omitempty" xml:"type"`

	// 服务类型。
	ServiceType *string `json:"service_type,omitempty" xml:"service_type"`

	// 状态。
	Status *string `json:"status,omitempty" xml:"status"`

	// 工作空间id。
	WorkspaceId *string `json:"workspace_id,omitempty" xml:"workspace_id"`

	// 创建时间。
	CreatedAt *int64 `json:"created_at,omitempty" xml:"created_at"`

	// 更新时间。
	UpdateAt *int64 `json:"update_at,omitempty" xml:"update_at"`

	SpecsConfig *SpecsConfig `json:"specs_config,omitempty" xml:"specs_config"`
}

func (o ResScene) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResScene struct{}"
	}

	return strings.Join([]string{"ResScene", string(data)}, " ")
}
