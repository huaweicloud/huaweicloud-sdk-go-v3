package model

import (
	"encoding/json"

	"strings"
)

type OrderLineItemV3 struct {
	// |参数名称：用于合并的订单项列表，会将相同产品、相同规格（对于线性产品）、相同最终价格（例如，严选产品改价）的进行合并| |参数约束以及描述：用于合并的订单项列表，会将相同产品、相同规格（对于线性产品）、相同最终价格（例如，严选产品改价）的进行合并|
	OrderLineItemIds []string `json:"order_line_item_ids"`
	// |参数名称：折扣模式 0：折扣 1：一口价 2：满减| |参数的约束及描述：折扣模式 0：折扣 1：一口价 2：满减|
	DiscountMode int32 `json:"discount_mode"`
	// |参数名称：折扣额（减免金额）| |参数的约束及描述：折扣额（减免金额）|
	DiscountAmount float32 `json:"discount_amount"`
	// |参数名称：折扣比例，折扣为| |参数的约束及描述：折扣比例，折扣为|
	DiscountRatio float32 `json:"discount_ratio"`
}

func (o OrderLineItemV3) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "OrderLineItemV3 struct{}"
	}

	return strings.Join([]string{"OrderLineItemV3", string(data)}, " ")
}
