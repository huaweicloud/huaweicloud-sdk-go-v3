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
type NeutronDeleteFirewallGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o NeutronDeleteFirewallGroupResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"NeutronDeleteFirewallGroupResponse", string(data)}, " ")
}
