package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubNodeLineageConfig 导入血缘请求体。
type SubNodeLineageConfig struct {

	// 节点名称。
	Name string `json:"name"`

	// 待接续sql语句。
	SqlStatement *string `json:"sql_statement,omitempty"`

	// 输入表。
	InputTables []TableConfig `json:"input_tables"`

	OutputTable *TableConfig `json:"output_table"`

	// 字段血缘信息。
	ColumnLineages *[]ColumnLineageConfig `json:"column_lineages,omitempty"`
}

func (o SubNodeLineageConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubNodeLineageConfig struct{}"
	}

	return strings.Join([]string{"SubNodeLineageConfig", string(data)}, " ")
}
