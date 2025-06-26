package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceSignPolicyExecTasksResponse Response Object
type ListInstanceSignPolicyExecTasksResponse struct {

	// 签名策略执行记录任务列表
	Tasks *[]SignatureExecutionTask `json:"tasks,omitempty"`

	// 签名策略执行记录任务总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstanceSignPolicyExecTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceSignPolicyExecTasksResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceSignPolicyExecTasksResponse", string(data)}, " ")
}
