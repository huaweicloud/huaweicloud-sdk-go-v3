/*
 * CloudPipeline
 *
 * devcloud pipeline api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowAgentStatusResponse struct {
	// Agent状态
	Status *string `json:"status,omitempty"`
	// AgentID
	AgentId *string `json:"agent_id,omitempty"`
}

func (o ShowAgentStatusResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowAgentStatusResponse", string(data)}, " ")
}
