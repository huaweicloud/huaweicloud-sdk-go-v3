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
type DeletePortRequest struct {
	PortId string `json:"port_id"`
}

func (o DeletePortRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeletePortRequest", string(data)}, " ")
}
