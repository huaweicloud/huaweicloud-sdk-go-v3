package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RetryScriptExecutionReq 重试脚本请求体。
type RetryScriptExecutionReq struct {

	// 脚本执行记录ID。
	RecordId *string `json:"record_id,omitempty"`

	// 脚本执行任务ID。
	TaskId *string `json:"task_id,omitempty"`
}

func (o RetryScriptExecutionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryScriptExecutionReq struct{}"
	}

	return strings.Join([]string{"RetryScriptExecutionReq", string(data)}, " ")
}
