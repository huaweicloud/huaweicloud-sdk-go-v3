package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDesktopPoolsScriptExecTasksResponse Response Object
type ShowDesktopPoolsScriptExecTasksResponse struct {

	// 脚本执行任务列表。
	ScriptExecutionTasks *[]ScriptExecutionTask `json:"script_execution_tasks,omitempty"`

	// 总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowDesktopPoolsScriptExecTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDesktopPoolsScriptExecTasksResponse struct{}"
	}

	return strings.Join([]string{"ShowDesktopPoolsScriptExecTasksResponse", string(data)}, " ")
}
