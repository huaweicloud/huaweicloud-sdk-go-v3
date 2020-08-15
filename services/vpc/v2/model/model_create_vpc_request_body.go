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
type CreateVpcRequestBody struct {
	Vpc *CreateVpcOption `json:"vpc"`
}

func (o CreateVpcRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateVpcRequestBody", string(data)}, " ")
}
