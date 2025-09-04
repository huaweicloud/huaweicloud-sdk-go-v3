package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportSessionsResponse Response Object
type ExportSessionsResponse struct {

	// 导出任务id。
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportSessionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportSessionsResponse struct{}"
	}

	return strings.Join([]string{"ExportSessionsResponse", string(data)}, " ")
}
