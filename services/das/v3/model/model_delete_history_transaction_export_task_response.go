package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHistoryTransactionExportTaskResponse Response Object
type DeleteHistoryTransactionExportTaskResponse struct {

	// 导出任务id
	TaskId         *int64 `json:"task_id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o DeleteHistoryTransactionExportTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHistoryTransactionExportTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteHistoryTransactionExportTaskResponse", string(data)}, " ")
}
