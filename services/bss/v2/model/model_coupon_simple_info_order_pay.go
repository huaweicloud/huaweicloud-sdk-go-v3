package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CouponSimpleInfoOrderPay struct {
	// 折扣ID。

	Id string `json:"id"`
	// 折扣类型： 0：促销折扣 1：合同折扣2：商务优惠3：合作伙伴授予折扣609：订单调价折扣 说明： 订单支付时，如果包含609折扣（订单调价折扣），则必须使用，不能再更换折扣类型。

	Type int32 `json:"type"`
}

func (o CouponSimpleInfoOrderPay) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CouponSimpleInfoOrderPay struct{}"
	}

	return strings.Join([]string{"CouponSimpleInfoOrderPay", string(data)}, " ")
}
