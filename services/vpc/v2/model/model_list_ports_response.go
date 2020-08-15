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
type ListPortsResponse struct {
	// port列表对象
	Ports []Port `json:"ports,omitempty"`
}

func (o ListPortsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListPortsResponse", string(data)}, " ")
}
