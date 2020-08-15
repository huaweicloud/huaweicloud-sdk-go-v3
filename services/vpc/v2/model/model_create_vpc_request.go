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

// Request Object
type CreateVpcRequest struct {
	Body *CreateVpcRequestBody `json:"body,omitempty"`
}

func (o CreateVpcRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateVpcRequest", string(data)}, " ")
}
