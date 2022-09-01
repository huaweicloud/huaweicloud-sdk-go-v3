package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type Datasource struct {

	// 数据源名称。
	DatasourceName string `json:"datasource_name" xml:"datasource_name"`

	// 工作空间id。
	WorkspaceId string `json:"workspace_id" xml:"workspace_id"`

	// 数据源id。
	DatasourceId *string `json:"datasource_id,omitempty" xml:"datasource_id"`

	// 状态。
	Status *string `json:"status,omitempty" xml:"status"`

	// 创建时间。
	CreatedAt *int64 `json:"created_at,omitempty" xml:"created_at"`
}

func (o Datasource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Datasource struct{}"
	}

	return strings.Join([]string{"Datasource", string(data)}, " ")
}
