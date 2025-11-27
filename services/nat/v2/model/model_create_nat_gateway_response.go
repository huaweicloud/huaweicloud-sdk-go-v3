package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateNatGatewayResponse Response Object
type CreateNatGatewayResponse struct {
	NatGateway *NatGatewayResponseBody `json:"nat_gateway,omitempty"`

	// 订单ID。
	OrderId *string `json:"order_id,omitempty"`

	// 公网NAT网关实例的ID。
	NatGatewayId   *string `json:"nat_gateway_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateNatGatewayResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNatGatewayResponse struct{}"
	}

	return strings.Join([]string{"CreateNatGatewayResponse", string(data)}, " ")
}
