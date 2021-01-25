package model

import (
	"encoding/json"

	"strings"
)

type OrderV3 struct {
	// |参数名称：订单标识| |参数约束及描述：订单标识|
	OrderId string `json:"order_id"`
	// |参数名称：订单行列表| |参数约束以及描述：订单行列表|
	OrderLineItems []OrderLineItemV3 `json:"order_line_items"`
}

func (o OrderV3) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "OrderV3 struct{}"
	}

	return strings.Join([]string{"OrderV3", string(data)}, " ")
}
