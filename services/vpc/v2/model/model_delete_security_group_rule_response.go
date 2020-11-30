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
type DeleteSecurityGroupRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSecurityGroupRuleResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteSecurityGroupRuleResponse", string(data)}, " ")
}
