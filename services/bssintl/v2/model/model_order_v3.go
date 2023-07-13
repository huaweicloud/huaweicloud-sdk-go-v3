package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OrderV3 struct {

	// 可使用折扣的订单ID。
	OrderId string `json:"order_id"`

	// 可使用折扣的订单项列表，具体参见表4。
	OrderLineItems []OrderLineItemV3 `json:"order_line_items"`
}

func (o OrderV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrderV3 struct{}"
	}

	return strings.Join([]string{"OrderV3", string(data)}, " ")
}
