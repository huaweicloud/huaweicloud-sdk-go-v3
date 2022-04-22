package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DiscountInfoV3 struct {

	// 订单的可用折扣ID。 支付订单时，输入该参数的值，即可使用折扣。
	DiscountId string `json:"discount_id"`

	// 折扣率或者满减值，如果折扣模式是一口价，这个值为空。
	DiscountValue string `json:"discount_value"`

	// 折扣类型，取值为 0：促销折扣1：合同折扣2：商务优惠3：合作伙伴授予折扣609：订单调价折扣
	DiscountType int32 `json:"discount_type"`

	// 可使用折扣的订单列表。 具体请参见表3。
	Orders []OrderV3 `json:"orders"`
}

func (o DiscountInfoV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiscountInfoV3 struct{}"
	}

	return strings.Join([]string{"DiscountInfoV3", string(data)}, " ")
}
