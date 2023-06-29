package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAgentRequest Request Object
type ListAgentRequest struct {

	// 每页显示条目数，正整数
	Limit *string `json:"limit,omitempty"`

	// 偏移值,正整数
	Offset *int32 `json:"offset,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 客户端ID
	AgentId *string `json:"agent_id,omitempty"`
}

func (o ListAgentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAgentRequest struct{}"
	}

	return strings.Join([]string{"ListAgentRequest", string(data)}, " ")
}
