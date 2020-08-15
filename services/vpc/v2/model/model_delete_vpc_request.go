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
type DeleteVpcRequest struct {
	VpcId string `json:"vpc_id"`
}

func (o DeleteVpcRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteVpcRequest", string(data)}, " ")
}
