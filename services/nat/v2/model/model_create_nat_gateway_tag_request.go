package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateNatGatewayTagRequest Request Object
type CreateNatGatewayTagRequest struct {

	// 所属公网NAT网关的id。
	NatGatewayId string `json:"nat_gateway_id"`

	Body *CreateNatTagRequestBody `json:"body,omitempty"`
}

func (o CreateNatGatewayTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNatGatewayTagRequest struct{}"
	}

	return strings.Join([]string{"CreateNatGatewayTagRequest", string(data)}, " ")
}
