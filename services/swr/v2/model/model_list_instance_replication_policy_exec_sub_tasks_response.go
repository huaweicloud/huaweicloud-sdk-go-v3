package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceReplicationPolicyExecSubTasksResponse Response Object
type ListInstanceReplicationPolicyExecSubTasksResponse struct {

	// 老化策略执行记录子任务列表
	Subtasks *[]SubtaskDetail `json:"subtasks,omitempty"`

	// 老化策略执行记录子任务总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstanceReplicationPolicyExecSubTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceReplicationPolicyExecSubTasksResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceReplicationPolicyExecSubTasksResponse", string(data)}, " ")
}
