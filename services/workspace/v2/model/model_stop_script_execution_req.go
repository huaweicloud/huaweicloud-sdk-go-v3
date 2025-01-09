package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopScriptExecutionReq 停止脚本或者命令执行任务请求体。
type StopScriptExecutionReq struct {

	// 脚本执行记录ID。
	RecordId *string `json:"record_id,omitempty"`

	// 脚本执行任务ID。
	TaskId *string `json:"task_id,omitempty"`
}

func (o StopScriptExecutionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopScriptExecutionReq struct{}"
	}

	return strings.Join([]string{"StopScriptExecutionReq", string(data)}, " ")
}
