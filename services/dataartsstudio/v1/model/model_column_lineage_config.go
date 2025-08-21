package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ColumnLineageConfig 导入血缘请求体。
type ColumnLineageConfig struct {

	// 字段血缘节点名称。
	Name *string `json:"name,omitempty"`

	// 输入字段血缘信息。
	InputColumns []ColumnConfig `json:"input_columns"`

	OutputColumn *ColumnConfig `json:"output_column"`
}

func (o ColumnLineageConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ColumnLineageConfig struct{}"
	}

	return strings.Join([]string{"ColumnLineageConfig", string(data)}, " ")
}
