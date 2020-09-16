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
type ListVpcsResponse struct {
	// vpc对象列表
	Vpcs *[]Vpc `json:"vpcs,omitempty"`
}

func (o ListVpcsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListVpcsResponse", string(data)}, " ")
}
