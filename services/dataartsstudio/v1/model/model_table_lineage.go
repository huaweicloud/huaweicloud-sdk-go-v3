package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TableLineage struct {

	// 作业算子名称
	Name string `json:"name"`

	// 上游血缘表列表，列表大小：1至100
	InputTables []TableInfo `json:"input_tables"`

	// 下游血缘表列表，列表大小：1至100
	OutputTables []TableInfo `json:"output_tables"`

	// 源数据连接id
	SourceConnectionId string `json:"source_connection_id"`

	// 目标数据连接id
	TargetConnectionId *string `json:"target_connection_id,omitempty"`

	// 字段血缘列表，列表大小：0至100
	ColumnLineages *[]ColumnLineage `json:"column_lineages,omitempty"`
}

func (o TableLineage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableLineage struct{}"
	}

	return strings.Join([]string{"TableLineage", string(data)}, " ")
}
