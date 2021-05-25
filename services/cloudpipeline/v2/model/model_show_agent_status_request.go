package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowAgentStatusRequest struct {
	// 语言类型 中文:zh-cn 英文:en-us

	XLanguage *string `json:"X-Language,omitempty"`
	// AgentID

	AgentId string `json:"agent_id"`
}

func (o ShowAgentStatusRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowAgentStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowAgentStatusRequest", string(data)}, " ")
}
