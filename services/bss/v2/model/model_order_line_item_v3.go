package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/shopspring/decimal"

	"strings"
)

type OrderLineItemV3 struct {

	// 可使用相同折扣的订单项合并后的订单项ID列表。 相同产品、相同规格（对于线性产品）、相同最终价格（例如，严选产品改价）的订单项将进行合并。
	OrderLineItemIds []string `json:"order_line_item_ids"`

	// 订单可用折扣的模式 。 0：折扣1：一口价2：满减
	DiscountMode int32 `json:"discount_mode"`

	// 订单可用的折扣金额（即减免金额）。
	DiscountAmount *decimal.Decimal `json:"discount_amount"`

	// 订单可用的折扣比例。
	DiscountRatio *decimal.Decimal `json:"discount_ratio"`
}

func (o OrderLineItemV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrderLineItemV3 struct{}"
	}

	return strings.Join([]string{"OrderLineItemV3", string(data)}, " ")
}
