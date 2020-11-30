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
type CreateVpcResponse struct {
	Vpc            *Vpc `json:"vpc,omitempty"`
	HttpStatusCode int  `json:"-"`
}

func (o CreateVpcResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateVpcResponse", string(data)}, " ")
}
