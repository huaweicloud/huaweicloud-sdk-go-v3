package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type NeutronShowFloatingIpRequest struct {
	// floatingip的ID

	FloatingipId string `json:"floatingip_id"`
}

func (o NeutronShowFloatingIpRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "NeutronShowFloatingIpRequest struct{}"
	}

	return strings.Join([]string{"NeutronShowFloatingIpRequest", string(data)}, " ")
}
