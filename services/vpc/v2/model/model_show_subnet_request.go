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
type ShowSubnetRequest struct {
	SubnetId string `json:"subnet_id"`
}

func (o ShowSubnetRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowSubnetRequest", string(data)}, " ")
}
