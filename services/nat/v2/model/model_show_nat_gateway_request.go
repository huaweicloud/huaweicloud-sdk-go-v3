package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNatGatewayRequest Request Object
type ShowNatGatewayRequest struct {

	// 公网NAT网关实例的ID。
	NatGatewayId string `json:"nat_gateway_id"`
}

func (o ShowNatGatewayRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNatGatewayRequest struct{}"
	}

	return strings.Join([]string{"ShowNatGatewayRequest", string(data)}, " ")
}
