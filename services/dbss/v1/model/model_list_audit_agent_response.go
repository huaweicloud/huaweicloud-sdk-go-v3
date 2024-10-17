package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditAgentResponse Response Object
type ListAuditAgentResponse struct {

	// agent列表
	Agents         *[]AuditAgentRespoonseAgents `json:"agents,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListAuditAgentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditAgentResponse struct{}"
	}

	return strings.Join([]string{"ListAuditAgentResponse", string(data)}, " ")
}
