package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type AssociateServerVirtualIpRequestBody struct {
	Nic *AssociateServerVirtualIpOption `json:"nic"`
}

func (o AssociateServerVirtualIpRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AssociateServerVirtualIpRequestBody struct{}"
	}

	return strings.Join([]string{"AssociateServerVirtualIpRequestBody", string(data)}, " ")
}
