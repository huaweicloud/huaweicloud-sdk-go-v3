package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOrderResponse Response Object
type CreateOrderResponse struct {

	// 订单号，下单成功时返回订单ID。
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateOrderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrderResponse struct{}"
	}

	return strings.Join([]string{"CreateOrderResponse", string(data)}, " ")
}
