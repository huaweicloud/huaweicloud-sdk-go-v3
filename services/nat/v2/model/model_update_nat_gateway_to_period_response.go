package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNatGatewayToPeriodResponse Response Object
type UpdateNatGatewayToPeriodResponse struct {

	// 订单ID。
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateNatGatewayToPeriodResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNatGatewayToPeriodResponse struct{}"
	}

	return strings.Join([]string{"UpdateNatGatewayToPeriodResponse", string(data)}, " ")
}
