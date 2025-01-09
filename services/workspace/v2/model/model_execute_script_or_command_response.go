package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteScriptOrCommandResponse Response Object
type ExecuteScriptOrCommandResponse struct {

	// 执行脚本的任务ID。
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExecuteScriptOrCommandResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteScriptOrCommandResponse struct{}"
	}

	return strings.Join([]string{"ExecuteScriptOrCommandResponse", string(data)}, " ")
}
