package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNatGatewayTagRequest Request Object
type ShowNatGatewayTagRequest struct {

	// 公网NAT网关ID。
	NatGatewayId string `json:"nat_gateway_id"`
}

func (o ShowNatGatewayTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNatGatewayTagRequest struct{}"
	}

	return strings.Join([]string{"ShowNatGatewayTagRequest", string(data)}, " ")
}
