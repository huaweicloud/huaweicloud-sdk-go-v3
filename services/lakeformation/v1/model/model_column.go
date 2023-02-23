package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 列信息
type Column struct {

	// 列类型，包括array bigint binary boolean char date decimal double float int interval map set smallint string struct timestamp tinyint union varchar
	ColumnType string `json:"column_type"`

	// 列名
	ColumnName string `json:"column_name"`

	// 列描述信息
	Comment *string `json:"comment,omitempty"`
}

func (o Column) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Column struct{}"
	}

	return strings.Join([]string{"Column", string(data)}, " ")
}
