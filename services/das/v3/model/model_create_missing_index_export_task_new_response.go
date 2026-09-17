package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMissingIndexExportTaskNewResponse Response Object
type CreateMissingIndexExportTaskNewResponse struct {

	// 任务ID
	TaskId         *int64 `json:"task_id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateMissingIndexExportTaskNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMissingIndexExportTaskNewResponse struct{}"
	}

	return strings.Join([]string{"CreateMissingIndexExportTaskNewResponse", string(data)}, " ")
}
