package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListIssuedPartnerCouponsRequest struct {
	// 优惠券ID。

	CouponId *string `json:"coupon_id,omitempty"`
	// 客户账号ID。您可以调用查询客户列表接口获取customer_id。

	CustomerId *string `json:"customer_id,omitempty"`
	// 订单ID。

	OrderId *string `json:"order_id,omitempty"`
	// 优惠券类型： 1：代金券4：现金券

	CouponType *int32 `json:"coupon_type,omitempty"`
	// 客户优惠券实例状态： 1：未激活2：可使用3：已使用4：已过期5：已回收

	Status *int32 `json:"status,omitempty"`
	// 创建时间（开始）。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。 其中，HH范围是0～23，mm和ss范围是0～59。 输入这个条件，会查询出创建时间大于这个时间的记录。

	CreateTimeBegin *string `json:"create_time_begin,omitempty"`
	// 创建时间（结束）。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。 其中，HH范围是0～23，mm和ss范围是0～59。 输入这个条件，会查询出创建时间小于这个时间的记录。

	CreateTimeEnd *string `json:"create_time_end,omitempty"`
	// 生效时间（开始）。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。 其中，HH范围是0～23，mm和ss范围是0～59。 输入这个条件，会查询出生效时间大于这个时间的记录。

	EffectiveTimeBegin *string `json:"effective_time_begin,omitempty"`
	// 生效时间（结束）。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。 其中，HH范围是0～23，mm和ss范围是0～59。 输入这个条件，会查询出生效时间小于这个时间的记录。

	EffectiveTimeEnd *string `json:"effective_time_end,omitempty"`
	// 失效时间（开始）。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。 其中，HH范围是0～23，mm和ss范围是0～59。 输入这个条件，会查询出失效时间大于这个时间的记录。

	ExpireTimeBegin *string `json:"expire_time_begin,omitempty"`
	// 失效时间（结束）。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。 其中，HH范围是0～23，mm和ss范围是0～59。 输入这个条件，会查询出失效时间小于这个时间的记录。

	ExpireTimeEnd *string `json:"expire_time_end,omitempty"`
	// 偏移量，从0开始。默认值为0。

	Offset *int32 `json:"offset,omitempty"`
	// 查询的每页数量。默认值为10。

	Limit *int32 `json:"limit,omitempty"`
	// 精英服务商ID。 华为云伙伴能力中心（一级经销商）查询精英服务商发放给子客户的优惠券时，需要携带该参数；否则只能查询发放给自己子客户的优惠券。

	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`
}

func (o ListIssuedPartnerCouponsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListIssuedPartnerCouponsRequest struct{}"
	}

	return strings.Join([]string{"ListIssuedPartnerCouponsRequest", string(data)}, " ")
}
