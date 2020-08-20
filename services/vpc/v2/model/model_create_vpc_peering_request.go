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
type CreateVpcPeeringRequest struct {
	Body *CreateVpcPeeringRequestBody `json:"body,omitempty"`
}

func (o CreateVpcPeeringRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateVpcPeeringRequest", string(data)}, " ")
}
