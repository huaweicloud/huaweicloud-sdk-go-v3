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
type NeutronUpdateFloatingIpResponse struct {
	Floatingip *PostAndPutFloatingIpResp `json:"floatingip,omitempty"`
}

func (o NeutronUpdateFloatingIpResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"NeutronUpdateFloatingIpResponse", string(data)}, " ")
}
