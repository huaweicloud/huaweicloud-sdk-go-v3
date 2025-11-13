package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHistoryTransactionExportTaskResponse Response Object
type ListHistoryTransactionExportTaskResponse struct {

	// 导出任务列表总大小
	Total *int32 `json:"total,omitempty"`

	// 导出任务对应记录列表
	TaskList       *[]ExportTaskDetailResp `json:"task_list,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListHistoryTransactionExportTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHistoryTransactionExportTaskResponse struct{}"
	}

	return strings.Join([]string{"ListHistoryTransactionExportTaskResponse", string(data)}, " ")
}
