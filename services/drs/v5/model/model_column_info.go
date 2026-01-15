package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ColumnInfo 列信息
type ColumnInfo struct {

	// 列名
	ColumnName *string `json:"column_name,omitempty"`

	// 列类型
	ColumnType *string `json:"column_type,omitempty"`

	// 列的键类型。
	KeyType *string `json:"key_type,omitempty"`

	// 列映射后的名称
	ColumnMappedName *string `json:"column_mapped_name,omitempty"`

	// 该列是否过滤
	Status *bool `json:"status,omitempty"`

	// 该列是否partitionKey
	Partition *bool `json:"partition,omitempty"`
}

func (o ColumnInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ColumnInfo struct{}"
	}

	return strings.Join([]string{"ColumnInfo", string(data)}, " ")
}
