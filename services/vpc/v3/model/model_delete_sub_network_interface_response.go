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
type DeleteSubNetworkInterfaceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSubNetworkInterfaceResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteSubNetworkInterfaceResponse", string(data)}, " ")
}
