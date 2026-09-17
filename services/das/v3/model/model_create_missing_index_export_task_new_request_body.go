package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMissingIndexExportTaskNewRequestBody 创建缺失索引导出任务请求体
type CreateMissingIndexExportTaskNewRequestBody struct {

	// 导出类型
	ExportType string `json:"export_type"`

	// 采集时间
	CollectTime int64 `json:"collect_time"`

	// 桶名
	BucketName string `json:"bucket_name"`

	// 过滤条件
	Conditions []ExportCondition `json:"conditions"`

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

func (o CreateMissingIndexExportTaskNewRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMissingIndexExportTaskNewRequestBody struct{}"
	}

	return strings.Join([]string{"CreateMissingIndexExportTaskNewRequestBody", string(data)}, " ")
}
