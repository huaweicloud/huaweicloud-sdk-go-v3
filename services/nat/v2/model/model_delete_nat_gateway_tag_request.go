package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteNatGatewayTagRequest Request Object
type DeleteNatGatewayTagRequest struct {

	// 公网NAT网关id。
	NatGatewayId string `json:"nat_gateway_id"`

	// 标签key。
	Key string `json:"key"`
}

func (o DeleteNatGatewayTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNatGatewayTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteNatGatewayTagRequest", string(data)}, " ")
}
