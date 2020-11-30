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
type UpdateSubnetResponse struct {
	Subnet         *SubnetResult `json:"subnet,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o UpdateSubnetResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateSubnetResponse", string(data)}, " ")
}
