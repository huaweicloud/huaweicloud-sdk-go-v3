package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAgentInvocationsResponse Response Object
type ListAgentInvocationsResponse struct {

	// 任务列表
	Invocations *[]InvocationInfo `json:"invocations,omitempty"`

	// 任务列表总量
	Count          *int64 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAgentInvocationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAgentInvocationsResponse struct{}"
	}

	return strings.Join([]string{"ListAgentInvocationsResponse", string(data)}, " ")
}
