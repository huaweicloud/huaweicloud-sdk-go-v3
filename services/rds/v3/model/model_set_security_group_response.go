/*
 * RDS
 *
 * API v3
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type SetSecurityGroupResponse struct {
	// 任务ID
	WorkflowId     *string `json:"workflowId,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetSecurityGroupResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"SetSecurityGroupResponse", string(data)}, " ")
}
