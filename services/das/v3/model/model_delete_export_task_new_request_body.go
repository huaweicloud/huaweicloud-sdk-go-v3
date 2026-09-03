package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteExportTaskNewRequestBody 删除binlog导出任务请求体
type DeleteExportTaskNewRequestBody struct {

	// binlog导出任务ID
	ExportTaskId int64 `json:"export_task_id"`
}

func (o DeleteExportTaskNewRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteExportTaskNewRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteExportTaskNewRequestBody", string(data)}, " ")
}
