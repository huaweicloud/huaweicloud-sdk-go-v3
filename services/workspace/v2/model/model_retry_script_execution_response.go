package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RetryScriptExecutionResponse Response Object
type RetryScriptExecutionResponse struct {

	// 执行脚本的任务ID。
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RetryScriptExecutionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryScriptExecutionResponse struct{}"
	}

	return strings.Join([]string{"RetryScriptExecutionResponse", string(data)}, " ")
}
