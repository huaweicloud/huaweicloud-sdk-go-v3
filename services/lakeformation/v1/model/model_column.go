package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Column 列信息
type Column struct {

	// 列类型，包括array bigint binary boolean char date decimal double float int interval map set smallint string struct timestamp tinyint union varchar
	ColumnType string `json:"column_type"`

	// 列名
	ColumnName string `json:"column_name"`

	// 列描述信息
	Comment *string `json:"comment,omitempty"`

	// 用户自定义类型名称，当且仅当column_type为struct，字段可填写。该字段存在值时，表示用户自定义的高阶类型
	AdvanceColumnType *string `json:"advance_column_type,omitempty"`
}

func (o Column) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Column struct{}"
	}

	return strings.Join([]string{"Column", string(data)}, " ")
}
