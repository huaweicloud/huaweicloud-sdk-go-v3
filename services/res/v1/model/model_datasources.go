package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Datasources struct {

	// 数据源id。
	DatasourceId *string `json:"datasource_id,omitempty"`

	// 名称。
	DatasourceName *string `json:"datasource_name,omitempty"`

	// 状态。
	Status *string `json:"status,omitempty"`

	// 结构。
	Structure *string `json:"structure,omitempty"`

	// 工作空间编号。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	DataConfig *DataConfig `json:"data_config,omitempty"`

	SpecsConfig *SpecsConfig `json:"specs_config,omitempty"`

	// 创建时间。
	CreatedAt *int64 `json:"created_at,omitempty"`

	// 更新时间。
	UpdateAt *int64 `json:"update_at,omitempty"`
}

func (o Datasources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Datasources struct{}"
	}

	return strings.Join([]string{"Datasources", string(data)}, " ")
}
