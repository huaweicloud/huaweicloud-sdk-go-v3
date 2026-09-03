package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchBinlogParseRequestBody 查看binlog解析详情请求体
type SearchBinlogParseRequestBody struct {

	// 解析任务ID
	TaskId int64 `json:"task_id"`

	// 开始时间（Unix timestamp），单位：毫秒
	StartTime *int64 `json:"start_time,omitempty"`

	// 结束时间（Unix timestamp），单位：毫秒
	EndTime *int64 `json:"end_time,omitempty"`

	// 数据库名称
	DbName *string `json:"db_name,omitempty"`

	// 表名称
	TableName *string `json:"table_name,omitempty"`

	// SQL类型列表。取值范围：insert、update、delete、ddl
	TypeList *[]string `json:"type_list,omitempty"`

	// 页码
	CurPage int32 `json:"cur_page"`

	// 每页记录数
	PerPage int32 `json:"per_page"`

	// 筛选条件列表
	ColumnList *[]FilterColumn `json:"column_list,omitempty"`
}

func (o SearchBinlogParseRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchBinlogParseRequestBody struct{}"
	}

	return strings.Join([]string{"SearchBinlogParseRequestBody", string(data)}, " ")
}
