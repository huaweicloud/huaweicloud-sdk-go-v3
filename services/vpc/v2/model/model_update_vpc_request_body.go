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

//
type UpdateVpcRequestBody struct {
	Vpc *UpdateVpcOption `json:"vpc"`
}

func (o UpdateVpcRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateVpcRequestBody", string(data)}, " ")
}
