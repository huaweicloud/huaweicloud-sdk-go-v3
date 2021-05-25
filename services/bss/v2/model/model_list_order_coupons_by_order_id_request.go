package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListOrderCouponsByOrderIdRequest struct {
	// 订单ID。

	OrderId string `json:"order_id"`
}

func (o ListOrderCouponsByOrderIdRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListOrderCouponsByOrderIdRequest struct{}"
	}

	return strings.Join([]string{"ListOrderCouponsByOrderIdRequest", string(data)}, " ")
}
