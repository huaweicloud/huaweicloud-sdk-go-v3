package model

import (
	"encoding/json"

	"strings"
)

type DiscountInfoV3 struct {
	// |参数名称：折扣ID，支付的时候，如果要使用折扣，需要将这个值填入| |参数约束及描述：折扣ID，支付的时候，如果要使用折扣，需要将这个值填入|
	DiscountId string `json:"discount_id"`
	// |参数名称：discountType为促销折扣或合作伙伴授予折扣时，为折扣率| |参数约束及描述：discountType为促销折扣或合作伙伴授予折扣时，为折扣率|
	DiscountValue string `json:"discount_value"`
	// |参数名称：折扣类型，取值为1：合同折扣（可以有多组）2：商务优惠（仅有一组）3：合作伙伴授予折扣（仅有一组）609：订单调价折扣| |参数的约束及描述：折扣类型，取值为1：合同折扣（可以有多组）2：商务优惠（仅有一组）3：合作伙伴授予折扣（仅有一组）609：订单调价折扣|
	DiscountType int32 `json:"discount_type"`
	// |参数名称：订单列表| |参数约束以及描述：订单列表|
	Orders []OrderV3 `json:"orders"`
}

func (o DiscountInfoV3) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DiscountInfoV3 struct{}"
	}

	return strings.Join([]string{"DiscountInfoV3", string(data)}, " ")
}
