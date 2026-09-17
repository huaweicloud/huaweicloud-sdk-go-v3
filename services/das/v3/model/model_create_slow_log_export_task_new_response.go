package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSlowLogExportTaskNewResponse Response Object
type CreateSlowLogExportTaskNewResponse struct {

	// 任务ID
	TaskId         *int64 `json:"task_id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateSlowLogExportTaskNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSlowLogExportTaskNewResponse struct{}"
	}

	return strings.Join([]string{"CreateSlowLogExportTaskNewResponse", string(data)}, " ")
}
