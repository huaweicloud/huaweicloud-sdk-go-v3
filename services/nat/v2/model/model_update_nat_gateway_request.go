package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateNatGatewayRequest struct {
	// 公网NAT网关实例的ID。

	NatGatewayId string `json:"nat_gateway_id"`

	Body *UpdateNatGatewayRequestBody `json:"body,omitempty"`
}

func (o UpdateNatGatewayRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateNatGatewayRequest struct{}"
	}

	return strings.Join([]string{"UpdateNatGatewayRequest", string(data)}, " ")
}
