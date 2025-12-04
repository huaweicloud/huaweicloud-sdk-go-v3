package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNatGatewayToPeriodRequest Request Object
type UpdateNatGatewayToPeriodRequest struct {

	// 公网NAT网关实例的ID。
	NatGatewayId string `json:"nat_gateway_id"`

	Body *UpdateNatGatewayToPeriodRequestBody `json:"body,omitempty"`
}

func (o UpdateNatGatewayToPeriodRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNatGatewayToPeriodRequest struct{}"
	}

	return strings.Join([]string{"UpdateNatGatewayToPeriodRequest", string(data)}, " ")
}
