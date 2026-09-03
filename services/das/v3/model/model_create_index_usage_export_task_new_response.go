package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIndexUsageExportTaskNewResponse Response Object
type CreateIndexUsageExportTaskNewResponse struct {

	// 任务ID
	TaskId         *int64 `json:"task_id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateIndexUsageExportTaskNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIndexUsageExportTaskNewResponse struct{}"
	}

	return strings.Join([]string{"CreateIndexUsageExportTaskNewResponse", string(data)}, " ")
}
