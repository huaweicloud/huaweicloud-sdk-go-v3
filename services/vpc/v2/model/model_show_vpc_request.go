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
type ShowVpcRequest struct {
	VpcId string `json:"vpc_id"`
}

func (o ShowVpcRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowVpcRequest", string(data)}, " ")
}
