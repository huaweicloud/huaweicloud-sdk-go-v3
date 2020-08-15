/*
 * VPC
 *
 * VPC Open API
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// Request Object
type CreateSecurityGroupRuleRequest struct {
	Body *CreateSecurityGroupRuleRequestBody `json:"body,omitempty"`
}

func (o CreateSecurityGroupRuleRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateSecurityGroupRuleRequest", string(data)}, " ")
}
