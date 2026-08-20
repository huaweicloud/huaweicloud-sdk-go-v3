package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteExportTaskRequest Request Object
type ExecuteExportTaskRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	// 导出任务ID
	JobId string `json:"job_id"`

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ExecuteExportTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteExportTaskRequest struct{}"
	}

	return strings.Join([]string{"ExecuteExportTaskRequest", string(data)}, " ")
}
