package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IQueryUserPartnerCouponsResultV2 struct {

	// 优惠券ID。
	CouponId *string `json:"coupon_id,omitempty" xml:"coupon_id"`

	// 优惠券状态： 1：未激活2：可使用3：已使用4：已过期5：已回收
	Status *int32 `json:"status,omitempty" xml:"status"`

	// 客户账号ID。
	CustomerId *string `json:"customer_id,omitempty" xml:"customer_id"`

	// 优惠券类别： 1：代金券4：现金券
	CouponType *int32 `json:"coupon_type,omitempty" xml:"coupon_type"`

	// 优惠券面额单位。 1：元。
	MeasureId *int32 `json:"measure_id,omitempty" xml:"measure_id"`

	// 优惠券面值。
	FaceValue *float64 `json:"face_value,omitempty" xml:"face_value"`

	// 生效时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。
	EffectiveTime *string `json:"effective_time,omitempty" xml:"effective_time"`

	// 失效时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。
	ExpireTime *string `json:"expire_time,omitempty" xml:"expire_time"`

	// 订单ID。
	OrderId *string `json:"order_id,omitempty" xml:"order_id"`

	// 促销计划ID。
	PromotionPlanId *string `json:"promotion_plan_id,omitempty" xml:"promotion_plan_id"`

	// 促销计划名称。
	PromotionPlanName *string `json:"promotion_plan_name,omitempty" xml:"promotion_plan_name"`

	// 促销计划描述。
	PromotionPlanDesc *string `json:"promotion_plan_desc,omitempty" xml:"promotion_plan_desc"`

	// 介质类型。 1：电子券2：纸质券
	MediaType *int32 `json:"media_type,omitempty" xml:"media_type"`

	// 获取方式。 1：线上领取2：线上兑换3：线上发放4：线下获取5：事件赠送
	FetchMethod *int32 `json:"fetch_method,omitempty" xml:"fetch_method"`

	// 优惠券限制。 具体请参见表3。
	UseLimits *[]ICouponUseLimitInfoV2 `json:"use_limits,omitempty" xml:"use_limits"`

	// 优惠券的激活时间。
	ActiveTime *string `json:"active_time,omitempty" xml:"active_time"`

	// 优惠券的使用时间。
	LastUsedTime *string `json:"last_used_time,omitempty" xml:"last_used_time"`

	// 促销活动ID。
	PromotionId *string `json:"promotion_id,omitempty" xml:"promotion_id"`

	// 优惠券的创建时间。
	CreateTime *string `json:"create_time,omitempty" xml:"create_time"`

	// 优惠券余额。
	Balance *float64 `json:"balance,omitempty" xml:"balance"`

	// 锁定优惠券的订单ID。 如果为老版本优惠券，该值为空。
	LockOrderId *string `json:"lock_order_id,omitempty" xml:"lock_order_id"`

	// 优惠券是否冻结。 0：否1：是 可用优惠券接口返回时不包括冻结状态的优惠券。
	IsFrozen *string `json:"is_frozen,omitempty" xml:"is_frozen"`
}

func (o IQueryUserPartnerCouponsResultV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IQueryUserPartnerCouponsResultV2 struct{}"
	}

	return strings.Join([]string{"IQueryUserPartnerCouponsResultV2", string(data)}, " ")
}
