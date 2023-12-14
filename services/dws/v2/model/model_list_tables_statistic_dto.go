package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTablesStatisticDto 表倾斜率或脏页率列表。
type ListTablesStatisticDto struct {

	// 数据库名称。
	DbName *string `json:"db_name,omitempty"`

	// schema名称。
	SchemaName *string `json:"schema_name,omitempty"`

	// 表名。
	TableName *string `json:"table_name,omitempty"`

	// 所属用户。
	TableOwner *string `json:"table_owner,omitempty"`

	// 表大小。
	TableSize *string `json:"table_size,omitempty"`

	// 表倾斜率。
	SkewRate *float64 `json:"skew_rate,omitempty"`

	// 脏页率。
	DirtyPageRate *float64 `json:"dirty_page_rate,omitempty"`
}

func (o ListTablesStatisticDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTablesStatisticDto struct{}"
	}

	return strings.Join([]string{"ListTablesStatisticDto", string(data)}, " ")
}
