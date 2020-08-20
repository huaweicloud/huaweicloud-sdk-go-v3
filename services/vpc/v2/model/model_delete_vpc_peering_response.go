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

// Response Object
type DeleteVpcPeeringResponse struct {
}

func (o DeleteVpcPeeringResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteVpcPeeringResponse", string(data)}, " ")
}
