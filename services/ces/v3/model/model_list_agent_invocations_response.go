package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAgentInvocationsResponse Response Object
type ListAgentInvocationsResponse struct {

	// **参数解释**: 任务列表 **取值范围**: 返回数组长度为[0,100]
	Invocations *[]InvocationInfo `json:"invocations,omitempty"`

	// **参数解释**: 任务列表总量 **取值范围**: 数字范围为[0,9999999999999]
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
