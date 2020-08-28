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
type UpdateSecurityGroupResponse struct {
	// 请求ID
	RequestId     *string            `json:"request_id,omitempty"`
	SecurityGroup *SecurityGroupInfo `json:"security_group,omitempty"`
}

func (o UpdateSecurityGroupResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateSecurityGroupResponse", string(data)}, " ")
}
