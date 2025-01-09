package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSubnetBandwidthChangeOrderResponse Response Object
type CreateSubnetBandwidthChangeOrderResponse struct {

	// 订单号，下单成功时返回订单ID。
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSubnetBandwidthChangeOrderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubnetBandwidthChangeOrderResponse struct{}"
	}

	return strings.Join([]string{"CreateSubnetBandwidthChangeOrderResponse", string(data)}, " ")
}
