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
type DeleteSubNetworkInterfaceRequest struct {
	SubNetworkInterfaceId string `json:"sub_network_interface_id"`
}

func (o DeleteSubNetworkInterfaceRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteSubNetworkInterfaceRequest", string(data)}, " ")
}
