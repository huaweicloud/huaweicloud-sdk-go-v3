package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type AssociateServerVirtualIpResponse struct {
	// 云服务器网卡ID。

	PortId         *string `json:"port_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AssociateServerVirtualIpResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AssociateServerVirtualIpResponse struct{}"
	}

	return strings.Join([]string{"AssociateServerVirtualIpResponse", string(data)}, " ")
}
