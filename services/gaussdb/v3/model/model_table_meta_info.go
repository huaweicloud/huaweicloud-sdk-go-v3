package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TableMetaInfo struct {

	// 列的数据类型
	ColumnType *string `json:"column_type,omitempty"`

	// 额外的信息，例如是否是自动递增列
	Extra *string `json:"extra,omitempty"`

	// 是否允许为NULL
	IsNullable *string `json:"is_nullable,omitempty"`

	// 列的默认值
	ColumnDefault *string `json:"column_default,omitempty"`

	// 是否是索引列
	ColumnKey *string `json:"column_key,omitempty"`

	// 列名
	ColumnName *string `json:"column_name,omitempty"`
}

func (o TableMetaInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableMetaInfo struct{}"
	}

	return strings.Join([]string{"TableMetaInfo", string(data)}, " ")
}
