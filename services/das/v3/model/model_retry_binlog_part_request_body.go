package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RetryBinlogPartRequestBody 重试binlog解析任务部分请求体
type RetryBinlogPartRequestBody struct {

	// binlog解析任务ID
	TaskId int64 `json:"task_id"`

	// binlog解析错误ID
	ErrorId *[]string `json:"error_id,omitempty"`

	// 数据库名称
	DbName *string `json:"db_name,omitempty"`

	// 数据库表名称
	TableName *string `json:"table_name,omitempty"`

	// 列表信息
	ColumnList *[]ColumnInfo `json:"column_list,omitempty"`
}

func (o RetryBinlogPartRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryBinlogPartRequestBody struct{}"
	}

	return strings.Join([]string{"RetryBinlogPartRequestBody", string(data)}, " ")
}
