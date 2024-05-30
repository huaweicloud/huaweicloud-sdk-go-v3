package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIssuedPartnerCouponsRequest Request Object
type ListIssuedPartnerCouponsRequest struct {

	// 优惠券ID。此参数不携带或携带值为空时，不作为筛选条件；携带值为null时，作为筛选条件；不支持携带值为空串。
	CouponId *string `json:"coupon_id,omitempty"`

	// 客户账号ID。您可以调用[查询客户列表](https://support.huaweicloud.com/api-bpconsole/mc_00021.html)接口获取customer_id。此参数不携带或携带值为空时，不作为筛选条件；携带值为空串或携带值为null时，作为筛选条件。
	CustomerId *string `json:"customer_id,omitempty"`

	// 订单ID。此参数不携带或携带值为空时，不作为筛选条件；携带值为null时，作为筛选条件；不支持携带值为空串。
	OrderId *string `json:"order_id,omitempty"`

	// 优惠券类型。1：代金券4：现金券此参数不携带或携带值为空或携带值为null时，不作为筛选条件。
	CouponType *int32 `json:"coupon_type,omitempty"`

	// 客户优惠券实例状态：1：未激活2：可使用3：已使用4：已过期5：已回收此参数不携带或携带值为空或携带值为null时，不作为筛选条件。
	Status *int32 `json:"status,omitempty"`

	// 创建时间（开始）。UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。其中，HH范围是0～23，mm和ss范围是0～59。输入这个条件，会查询出创建时间大于这个时间的记录。此参数不携带或携带值为空时，不作为筛选条件。
	CreateTimeBegin *string `json:"create_time_begin,omitempty"`

	// 创建时间（结束）。UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。其中，HH范围是0～23，mm和ss范围是0～59。输入这个条件，会查询出创建时间小于这个时间的记录。此参数不携带或携带值为空时，不作为筛选条件。
	CreateTimeEnd *string `json:"create_time_end,omitempty"`

	// 生效时间（开始）。UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。其中，HH范围是0～23，mm和ss范围是0～59。输入这个条件，会查询出生效时间大于这个时间的记录。此参数不携带或携带值为空时，不作为筛选条件。
	EffectiveTimeBegin *string `json:"effective_time_begin,omitempty"`

	// 生效时间（结束）。UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。其中，HH范围是0～23，mm和ss范围是0～59。输入这个条件，会查询出生效时间小于这个时间的记录。此参数不携带或携带值为空时，不作为筛选条件。
	EffectiveTimeEnd *string `json:"effective_time_end,omitempty"`

	// 失效时间（开始）。UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。其中，HH范围是0～23，mm和ss范围是0～59。输入这个条件，会查询出失效时间大于这个时间的记录。此参数不携带或携带值为空时，不作为筛选条件。
	ExpireTimeBegin *string `json:"expire_time_begin,omitempty"`

	// 失效时间（结束）。UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。其中，HH范围是0～23，mm和ss范围是0～59。输入这个条件，会查询出失效时间小于这个时间的记录。此参数不携带或携带值为空时，不作为筛选条件。
	ExpireTimeEnd *string `json:"expire_time_end,omitempty"`

	// 偏移量，从0开始。默认值为0。 说明： offset用于分页处理，如不涉及分页，请使用默认值0。offset表示相对于满足条件的第一个数据的偏移量。如offset = 1，则返回满足条件的第二个数据至最后一个数据。例如，满足查询条件的结果共10条数据，limit取值为10，offset取值为1，则返回的数据为2~10，第一条数据不返回。
	Offset *int32 `json:"offset,omitempty"`

	// 查询的每页数量。默认值为10。
	Limit *int32 `json:"limit,omitempty"`

	// 云经销商ID。获取方法请参见[查询云经销商列表](https://support.huaweicloud.com/api-bpconsole/espp_00003.html)。华为云总经销商（一级经销商）查询云经销商发放给子客户的优惠券时，需要携带该参数；否则只能查询发放给自己子客户的优惠券。
	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`
}

func (o ListIssuedPartnerCouponsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIssuedPartnerCouponsRequest struct{}"
	}

	return strings.Join([]string{"ListIssuedPartnerCouponsRequest", string(data)}, " ")
}
