package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAgentStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowAgentStatusRequest", string(data)}, " ")
}
