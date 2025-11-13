package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateHistoryTransactionExportTaskResponse Response Object
type CreateHistoryTransactionExportTaskResponse struct {

	// 导出任务id
	TaskId         *int64 `json:"task_id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateHistoryTransactionExportTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHistoryTransactionExportTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateHistoryTransactionExportTaskResponse", string(data)}, " ")
}
