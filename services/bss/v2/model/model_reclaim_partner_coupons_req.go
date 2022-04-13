package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReclaimPartnerCouponsReq struct {
	// 待回收的代金券ID。 请从“发放优惠券”或“查询已发放的优惠券”接口的响应参数中获取。

	CouponId string `json:"coupon_id"`
	// 精英服务商ID。获取方法请参见查询精英服务商列表。 精英服务商回收给子客户发放的优惠券时，需要携带该字段。

	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`
}

func (o ReclaimPartnerCouponsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReclaimPartnerCouponsReq struct{}"
	}

	return strings.Join([]string{"ReclaimPartnerCouponsReq", string(data)}, " ")
}
