package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Datasources struct {

	// 数据源id。
	DatasourceId *string `json:"datasource_id,omitempty" xml:"datasource_id"`

	// 名称。
	DatasourceName *string `json:"datasource_name,omitempty" xml:"datasource_name"`

	// 状态。
	Status *string `json:"status,omitempty" xml:"status"`

	// 结构。
	Structure *string `json:"structure,omitempty" xml:"structure"`

	// 工作空间编号。
	WorkspaceId *string `json:"workspace_id,omitempty" xml:"workspace_id"`

	DataConfig *DataConfig `json:"data_config,omitempty" xml:"data_config"`

	SpecsConfig *SpecsConfig `json:"specs_config,omitempty" xml:"specs_config"`

	// 创建时间。
	CreatedAt *int64 `json:"created_at,omitempty" xml:"created_at"`

	// 更新时间。
	UpdateAt *int64 `json:"update_at,omitempty" xml:"update_at"`
}

func (o Datasources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Datasources struct{}"
	}

	return strings.Join([]string{"Datasources", string(data)}, " ")
}
