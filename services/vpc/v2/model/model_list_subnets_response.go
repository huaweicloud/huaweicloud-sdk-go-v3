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
type ListSubnetsResponse struct {
	// subnet对象列表
	Subnets        *[]Subnet `json:"subnets,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListSubnetsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListSubnetsResponse", string(data)}, " ")
}
