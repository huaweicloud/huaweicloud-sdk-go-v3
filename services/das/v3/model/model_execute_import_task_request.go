package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteImportTaskRequest Request Object
type ExecuteImportTaskRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	// 导入任务ID
	JobId string `json:"job_id"`

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ExecuteImportTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteImportTaskRequest struct{}"
	}

	return strings.Join([]string{"ExecuteImportTaskRequest", string(data)}, " ")
}
