package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportApiDefinitionsAsyncResponse Response Object
type ExportApiDefinitionsAsyncResponse struct {

	// 任务id
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportApiDefinitionsAsyncResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportApiDefinitionsAsyncResponse struct{}"
	}

	return strings.Join([]string{"ExportApiDefinitionsAsyncResponse", string(data)}, " ")
}
