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
type ShowVpcResponse struct {
	Vpc *Vpc `json:"vpc,omitempty"`
}

func (o ShowVpcResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowVpcResponse", string(data)}, " ")
}
