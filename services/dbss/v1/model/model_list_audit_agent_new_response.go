package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditAgentNewResponse Response Object
type ListAuditAgentNewResponse struct {

	// agent列表
	Agents         *[]AuditAgentRespoonseAgents `json:"agents,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListAuditAgentNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditAgentNewResponse struct{}"
	}

	return strings.Join([]string{"ListAuditAgentNewResponse", string(data)}, " ")
}
