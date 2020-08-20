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

// Response Object
type ShowSecurityGroupResponse struct {
	SecurityGroup *SecurityGroup `json:"security_group,omitempty"`
}

func (o ShowSecurityGroupResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowSecurityGroupResponse", string(data)}, " ")
}
