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

// Response Object
type NeutronCreateFloatingIpResponse struct {
	Floatingip *PostAndPutFloatingIpResp `json:"floatingip,omitempty"`
}

func (o NeutronCreateFloatingIpResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"NeutronCreateFloatingIpResponse", string(data)}, " ")
}
