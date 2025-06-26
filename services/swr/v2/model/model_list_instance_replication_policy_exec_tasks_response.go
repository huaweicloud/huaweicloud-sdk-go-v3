package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceReplicationPolicyExecTasksResponse Response Object
type ListInstanceReplicationPolicyExecTasksResponse struct {

	// 同步策略执行记录任务列表
	Tasks *[]TaskDetail `json:"tasks,omitempty"`

	// 同步策略执行记录任务总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstanceReplicationPolicyExecTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceReplicationPolicyExecTasksResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceReplicationPolicyExecTasksResponse", string(data)}, " ")
}
