/*
 * EIP
 *
 * 云服务接口
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// Request Object
type NeutronCreateFloatingIpRequest struct {
	Body *NeutronCreateFloatingIpRequestBody `json:"body,omitempty"`
}

func (o NeutronCreateFloatingIpRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"NeutronCreateFloatingIpRequest", string(data)}, " ")
}
