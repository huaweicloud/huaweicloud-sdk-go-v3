package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIndexUsageExportTaskNewRequestBody 创建索引使用导出任务请求体
type CreateIndexUsageExportTaskNewRequestBody struct {

	// 导出类型。取值范围：missingindex（导出表数据）、missingindexscript（导出脚本）
	ExportType string `json:"export_type"`

	// 采集时间
	CollectTime int64 `json:"collect_time"`

	// 桶名
	BucketName string `json:"bucket_name"`

	// 过滤条件
	Conditions *[]IndexUsageCondition `json:"conditions,omitempty"`

	// 表名称
	ObjectName *string `json:"object_name,omitempty"`

	// 排序字段
	SortField *string `json:"sort_field,omitempty"`

	// 排序是否升序
	SortAsc *bool `json:"sort_asc,omitempty"`

	// 当前页
	CurPage *int32 `json:"cur_page,omitempty"`

	// 页大小
	PerPage *int32 `json:"per_page,omitempty"`
}

func (o CreateIndexUsageExportTaskNewRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIndexUsageExportTaskNewRequestBody struct{}"
	}

	return strings.Join([]string{"CreateIndexUsageExportTaskNewRequestBody", string(data)}, " ")
}
