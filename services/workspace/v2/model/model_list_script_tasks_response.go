package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScriptTasksResponse Response Object
type ListScriptTasksResponse struct {

	// 总数。
	Count *int32 `json:"count,omitempty"`

	// 脚本任务列表。
	ScriptTasks    *[]ScriptTaskInfo `json:"script_tasks,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListScriptTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScriptTasksResponse struct{}"
	}

	return strings.Join([]string{"ListScriptTasksResponse", string(data)}, " ")
}
