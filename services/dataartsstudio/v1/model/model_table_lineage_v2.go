package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TableLineageV2 struct {

	// 上游血缘表列表，列表大小：1至100
	InputTables []TableInfoV2 `json:"input_tables"`

	// 下游血缘表列表，列表大小：1至100
	OutputTables []TableInfoV2 `json:"output_tables"`

	// 字段血缘列表，列表大小：0至100
	ColumnLineages *[]ColumnLineageV2 `json:"column_lineages,omitempty"`
}

func (o TableLineageV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableLineageV2 struct{}"
	}

	return strings.Join([]string{"TableLineageV2", string(data)}, " ")
}
