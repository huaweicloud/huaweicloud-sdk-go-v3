package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListSubCustomerCouponsRequest struct {

	// 优惠券ID。此参数不携带或携带值为空时，不作为筛选条件；携带值为空串时，作为筛选条件。
	CouponId *string `json:"coupon_id,omitempty" xml:"coupon_id"`

	// 订单ID。此参数不携带或携带值为空时，不作为筛选条件；携带值为空串时，作为筛选条件。
	OrderId *string `json:"order_id,omitempty" xml:"order_id"`

	// 促销计划ID。此参数不携带或携带值为空时，不作为筛选条件；携带值为空串时，作为筛选条件。
	PromotionPlanId *string `json:"promotion_plan_id,omitempty" xml:"promotion_plan_id"`

	// 优惠券类型：1：代金券2：折扣券3：产品券4：现金券此参数不携带或携带值为空时，不作为筛选条件。
	CouponType *int32 `json:"coupon_type,omitempty" xml:"coupon_type"`

	// 客户优惠券实例状态：1：未激活2：待使用3：已使用4：已过期此参数不携带或携带值为空时，不作为筛选条件。
	Status *int32 `json:"status,omitempty" xml:"status"`

	// 激活时间。UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。其中，HH范围是0～23，mm和ss范围是0～59。此参数不携带或携带值为空时，不作为筛选条件。
	ActiveStartTime *string `json:"active_start_time,omitempty" xml:"active_start_time"`

	// 结束时间。UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。其中，HH范围是0～23，mm和ss范围是0～59。此参数不携带或携带值为空时，不作为筛选条件。
	ActiveEndTime *string `json:"active_end_time,omitempty" xml:"active_end_time"`

	// 偏移量，从0开始。默认值为0。 说明： offset用于分页处理，如不涉及分页，请使用默认值0。offset表示相对于满足条件的第一个数据的偏移量。如offset = 1，则返回满足条件的第二个数据至最后一个数据。例如，满足查询条件的结果共10条数据，limit取值为10，offset取值为1，则返回的数据为2~10，第一条数据不返回。
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 查询的优惠券数量，默认值为10。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 发券来源，如果是合作伙伴发送的券，此处为伙伴ID。如果需要查询某个伙伴发放的券，可以在此处输入该伙伴ID。此参数不携带或携带值为空时，不作为筛选条件；携带值为空串时，作为筛选条件。
	SourceId *string `json:"source_id,omitempty" xml:"source_id"`

	// 云经销商（二级经销商）ID。华为云总经销商（一级经销商）查询云经销商名下的优惠券时，需要携带该参数；除此之外，此参数不做处理。否则只能查询自己的优惠券列表。
	IndirectPartnerId *string `json:"indirect_partner_id,omitempty" xml:"indirect_partner_id"`
}

func (o ListSubCustomerCouponsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubCustomerCouponsRequest struct{}"
	}

	return strings.Join([]string{"ListSubCustomerCouponsRequest", string(data)}, " ")
}
