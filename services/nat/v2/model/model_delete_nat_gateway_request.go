package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteNatGatewayRequest struct {
	// 公网NAT网关实例的ID。

	NatGatewayId string `json:"nat_gateway_id"`
}

func (o DeleteNatGatewayRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteNatGatewayRequest struct{}"
	}

	return strings.Join([]string{"DeleteNatGatewayRequest", string(data)}, " ")
}
