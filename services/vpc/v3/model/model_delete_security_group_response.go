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
type DeleteSecurityGroupResponse struct {
}

func (o DeleteSecurityGroupResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteSecurityGroupResponse", string(data)}, " ")
}
