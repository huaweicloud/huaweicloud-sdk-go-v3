package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddColumnInfo struct {

	// 列类型 取值：default_value，create_time，update_time，expression，server_database_table
	ColumnType *string `json:"column_type,omitempty"`

	// 列名称
	ColumnName *string `json:"column_name,omitempty"`

	// 列填充值
	ColumnValue *string `json:"column_value,omitempty"`

	// 填充列的数据类型 取值：int，long，varchar(256)，varchar(191)，datetime，timestamp
	DataType *string `json:"data_type,omitempty"`
}

func (o AddColumnInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddColumnInfo struct{}"
	}

	return strings.Join([]string{"AddColumnInfo", string(data)}, " ")
}
