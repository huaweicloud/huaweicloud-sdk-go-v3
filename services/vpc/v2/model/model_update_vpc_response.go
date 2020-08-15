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
type UpdateVpcResponse struct {
	Vpc *Vpc `json:"vpc,omitempty"`
}

func (o UpdateVpcResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateVpcResponse", string(data)}, " ")
}
