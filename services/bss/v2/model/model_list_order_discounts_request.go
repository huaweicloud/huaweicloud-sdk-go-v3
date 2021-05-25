package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListOrderDiscountsRequest struct {
	// 订单ID。

	OrderId string `json:"order_id"`
}

func (o ListOrderDiscountsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListOrderDiscountsRequest struct{}"
	}

	return strings.Join([]string{"ListOrderDiscountsRequest", string(data)}, " ")
}
