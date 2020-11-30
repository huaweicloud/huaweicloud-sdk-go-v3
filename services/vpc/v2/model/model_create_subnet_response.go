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
type CreateSubnetResponse struct {
	Subnet         *Subnet `json:"subnet,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSubnetResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateSubnetResponse", string(data)}, " ")
}
