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
type ListSubnetsRequest struct {
	Limit  int32  `json:"limit,omitempty"`
	Marker string `json:"marker,omitempty"`
	VpcId  string `json:"vpc_id,omitempty"`
}

func (o ListSubnetsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListSubnetsRequest", string(data)}, " ")
}
