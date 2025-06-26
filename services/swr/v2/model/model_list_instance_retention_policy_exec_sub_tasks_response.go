package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceRetentionPolicyExecSubTasksResponse Response Object
type ListInstanceRetentionPolicyExecSubTasksResponse struct {

	// 老化策略执行记录子任务列表
	SubTasks *[]Subtask `json:"sub_tasks,omitempty"`

	// 老化策略执行记录子任务总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstanceRetentionPolicyExecSubTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceRetentionPolicyExecSubTasksResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceRetentionPolicyExecSubTasksResponse", string(data)}, " ")
}
