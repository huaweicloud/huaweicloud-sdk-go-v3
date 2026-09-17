package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSlowLogExportTaskNewRequestBody 创建慢日志导出任务请求体
type CreateSlowLogExportTaskNewRequestBody struct {

	// 开始时间（Unix timestamp），单位：毫秒
	StartTime int64 `json:"start_time"`

	// 结束时间（Unix timestamp），单位：毫秒
	EndTime int64 `json:"end_time"`

	// OBS桶名
	BucketName string `json:"bucket_name"`

	// 文件目录
	FilePath *string `json:"file_path,omitempty"`

	// 导出类型
	ExportType *string `json:"export_type,omitempty"`

	// 排序字段
	SortField *string `json:"sort_field,omitempty"`

	// 排序顺序（true：正序，false：逆序）
	SortAsc *bool `json:"sort_asc,omitempty"`

	// 客户端
	Client *string `json:"client,omitempty"`

	// 用户
	User *string `json:"user,omitempty"`

	// 执行状态
	Killed *string `json:"killed,omitempty"`

	// 最小执行时间（Unix timestamp），单位：毫秒
	ExecuteTimeMin *int64 `json:"execute_time_min,omitempty"`

	// 最大执行时间（Unix timestamp），单位：毫秒
	ExecuteTimeMax *int64 `json:"execute_time_max,omitempty"`

	// 最小平均执行时间
	MinAvgExecuteTime *float64 `json:"min_avg_execute_time,omitempty"`

	// 最大平均执行时间
	MaxAvgExecuteTime *float64 `json:"max_avg_execute_time,omitempty"`

	// 最大扫描行数
	RowsMaxExamined *int64 `json:"rows_max_examined,omitempty"`

	// 最小扫描行数
	RowsMinExamined *int64 `json:"rows_min_examined,omitempty"`

	// 模糊SQL
	FuzzySql *string `json:"fuzzy_sql,omitempty"`

	// 操作（可组合，用逗号分隔）
	Operation *string `json:"operation,omitempty"`

	// 时区
	TimeZone *string `json:"time_zone,omitempty"`
}

func (o CreateSlowLogExportTaskNewRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSlowLogExportTaskNewRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSlowLogExportTaskNewRequestBody", string(data)}, " ")
}
