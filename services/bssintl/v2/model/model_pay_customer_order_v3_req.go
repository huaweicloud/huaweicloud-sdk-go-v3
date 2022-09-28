package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PayCustomerOrderV3Req struct {

	// 订单编号。 取值为调用“查询订单列表”接口时响应消息中的“order_id”字段的值或调用“续订包年/包月资源”接口时响应消息“order_ids”中的订单ID。
	OrderId string `json:"order_id"`

	// 本次订单支付是否使用优惠券。传递“YES”时，coupon_infos字段必选，传递“NO”时，会忽略coupon_infos字段的传值。 使用优惠券：YES，不使用优惠券：NO
	UseCoupon string `json:"use_coupon"`

	// 本次订单支付是否使用折扣。传递“YES”时，discount_infos字段必选，传递“NO”时，会忽略discount_infos字段的传值。 使用折扣：YES，不使用折扣：NO
	UseDiscount string `json:"use_discount"`

	// 优惠券ID列表，目前支持传递最多三个优惠券ID。 请从“查询订单可用优惠券”接口的响应参数中获取。 具体参见表1。 当use_coupon参数取值为“YES”，本字段必填；当use_coupon参数取值为“NO”，本字段不可填写，否则报参数错误。
	CouponInfos *[]CouponSimpleInfoOrderPayV3 `json:"coupon_infos,omitempty"`

	// 折扣ID列表，目前仅支持传递一个折扣ID。 请从“查询订单可用折扣”接口的响应参数中获取。 具体参见表2。 当use_discount参数取值为“YES”，本字段必填；当use_discount参数取值为“NO”，本字段不可填写，否则报参数错误。
	DiscountInfos *[]DiscountSimpleInfoV3 `json:"discount_infos,omitempty"`
}

func (o PayCustomerOrderV3Req) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PayCustomerOrderV3Req struct{}"
	}

	return strings.Join([]string{"PayCustomerOrderV3Req", string(data)}, " ")
}
