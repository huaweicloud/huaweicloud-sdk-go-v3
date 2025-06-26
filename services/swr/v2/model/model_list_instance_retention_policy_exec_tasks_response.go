package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceRetentionPolicyExecTasksResponse Response Object
type ListInstanceRetentionPolicyExecTasksResponse struct {

	// 老化策略执行记录任务列表
	Tasks *[]Task `json:"tasks,omitempty"`

	// 老化策略执行记录任务总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstanceRetentionPolicyExecTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceRetentionPolicyExecTasksResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceRetentionPolicyExecTasksResponse", string(data)}, " ")
}
