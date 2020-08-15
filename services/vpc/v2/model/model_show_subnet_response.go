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
type ShowSubnetResponse struct {
	Subnet *Subnet `json:"subnet,omitempty"`
}

func (o ShowSubnetResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowSubnetResponse", string(data)}, " ")
}
