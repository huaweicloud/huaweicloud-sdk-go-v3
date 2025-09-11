package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateAgentInvocationsResponse Response Object
type BatchCreateAgentInvocationsResponse struct {

	// **参数解释**: 创建任务的信息列表 **取值范围**: 数组长度范围为[0,100]
	Invocations    *[]BatchCreateInvocationInfo `json:"invocations,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o BatchCreateAgentInvocationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateAgentInvocationsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateAgentInvocationsResponse", string(data)}, " ")
}
