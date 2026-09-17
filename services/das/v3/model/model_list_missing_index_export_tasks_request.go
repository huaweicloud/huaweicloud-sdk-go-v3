package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMissingIndexExportTasksRequest Request Object
type ListMissingIndexExportTasksRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 导出类型，取值范围：missingindex（导出表数据）、missingindexscript（导出脚本）
	ExportType string `json:"export_type"`

	// 当前页
	CurPage *int32 `json:"cur_page,omitempty"`

	// 页大小
	PerPage *int32 `json:"per_page,omitempty"`
}

func (o ListMissingIndexExportTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMissingIndexExportTasksRequest struct{}"
	}

	return strings.Join([]string{"ListMissingIndexExportTasksRequest", string(data)}, " ")
}
