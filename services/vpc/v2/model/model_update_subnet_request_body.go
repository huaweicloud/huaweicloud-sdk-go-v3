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
type UpdateSubnetRequestBody struct {
	Subnet *UpdateSubnetOption `json:"subnet"`
}

func (o UpdateSubnetRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateSubnetRequestBody", string(data)}, " ")
}
