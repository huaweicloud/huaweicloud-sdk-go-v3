package model

import (
	"encoding/json"

	"strings"
)

type CancelCustomerOrderReq struct {
	// 订单ID。 取值为调用“查询订单列表”接口时响应消息中的“order_id”字段的值。

	OrderId string `json:"order_id"`
}

func (o CancelCustomerOrderReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CancelCustomerOrderReq struct{}"
	}

	return strings.Join([]string{"CancelCustomerOrderReq", string(data)}, " ")
}
