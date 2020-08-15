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
type UpdateVpcRequest struct {
	VpcId string                `json:"vpc_id"`
	Body  *UpdateVpcRequestBody `json:"body,omitempty"`
}

func (o UpdateVpcRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateVpcRequest", string(data)}, " ")
}
