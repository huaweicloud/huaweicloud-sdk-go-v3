package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Datasource
type Datasource struct {

	// 数据源名称。
	DatasourceName string `json:"datasource_name"`

	// 工作空间id。
	WorkspaceId string `json:"workspace_id"`

	// 数据源id。
	DatasourceId *string `json:"datasource_id,omitempty"`

	// 状态。
	Status *string `json:"status,omitempty"`

	// 创建时间。
	CreatedAt *int64 `json:"created_at,omitempty"`
}

func (o Datasource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Datasource struct{}"
	}

	return strings.Join([]string{"Datasource", string(data)}, " ")
}
