package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryColumnInfo 数据库列信息
type QueryColumnInfo struct {

	// 列名
	ColumnName *string `json:"column_name,omitempty"`

	// 列类型
	ColumnType *string `json:"column_type,omitempty"`

	// 主键或者唯一索引
	PrimaryKeyOrUniqueIndex *string `json:"primary_key_or_unique_index,omitempty"`

	// 列映射后的名称
	ColumnMappedName *string `json:"column_mapped_name,omitempty"`

	// 该列是否过滤
	IsFiltered *bool `json:"is_filtered,omitempty"`

	// 该列是否partitionKey
	IsPartitionKey *bool `json:"is_partition_key,omitempty"`
}

func (o QueryColumnInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryColumnInfo struct{}"
	}

	return strings.Join([]string{"QueryColumnInfo", string(data)}, " ")
}
