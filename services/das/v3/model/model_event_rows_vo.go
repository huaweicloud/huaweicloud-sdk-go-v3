package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventRowsVo binlog事件详情
type EventRowsVo struct {

	// 事件ID
	EventId *int64 `json:"event_id,omitempty"`

	// 文件名称
	FileName *string `json:"file_name,omitempty"`

	// 解析位置
	Position *int64 `json:"position,omitempty"`

	// 事件发生时间，单位毫秒
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 变更的数据库名称
	DbName *string `json:"db_name,omitempty"`

	// 变更的表名称
	TableName *string `json:"table_name,omitempty"`

	// 变更的SQL类型。取值范围：insert、update、delete、ddl
	SqlType *string `json:"sql_type,omitempty"`

	// 变更的SQL语句
	SqlStatement *string `json:"sql_statement,omitempty"`

	// 变更的列名称列表
	ColumnNames *[]string `json:"column_names,omitempty"`

	// 变更的主键列表
	PrimaryKeys *[]string `json:"primary_keys,omitempty"`

	// 变更影响的行数
	AffectRows *int32 `json:"affect_rows,omitempty"`

	// 变更数据详情
	RowPairs *[]RowPairDto `json:"row_pairs,omitempty"`
}

func (o EventRowsVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventRowsVo struct{}"
	}

	return strings.Join([]string{"EventRowsVo", string(data)}, " ")
}
