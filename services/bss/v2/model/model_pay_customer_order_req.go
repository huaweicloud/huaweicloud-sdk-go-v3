package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PayCustomerOrderReq struct {
	// 订单编号。 取值为调用“查询订单列表”接口时响应消息中的“order_id”字段的值或调用“续订包年/包月资源”接口时响应消息“order_ids”中的订单ID。

	OrderId string `json:"order_id"`
	// 优惠券ID列表，目前仅支持传递一个优惠券ID。 请从“查询订单可用优惠券”接口的响应参数中获取。 具体参见表1。

	CouponInfos *[]CouponSimpleInfoOrderPay `json:"coupon_infos,omitempty"`
	// 折扣ID列表，目前仅支持传递一个折扣ID。 请从“查询订单可用折扣”接口的响应参数中获取。 具体参见表2。

	DiscountInfos *[]DiscountSimpleInfo `json:"discount_infos,omitempty"`
}

func (o PayCustomerOrderReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PayCustomerOrderReq struct{}"
	}

	return strings.Join([]string{"PayCustomerOrderReq", string(data)}, " ")
}
