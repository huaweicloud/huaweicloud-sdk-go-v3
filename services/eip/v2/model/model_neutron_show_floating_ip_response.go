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
type NeutronShowFloatingIpResponse struct {
	Floatingip *FloatingIpResp `json:"floatingip,omitempty"`
}

func (o NeutronShowFloatingIpResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"NeutronShowFloatingIpResponse", string(data)}, " ")
}
