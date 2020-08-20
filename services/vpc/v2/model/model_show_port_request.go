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
type ShowPortRequest struct {
	PortId string `json:"port_id"`
}

func (o ShowPortRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowPortRequest", string(data)}, " ")
}
