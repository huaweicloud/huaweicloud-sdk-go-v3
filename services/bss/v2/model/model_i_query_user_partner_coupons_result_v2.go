package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/shopspring/decimal"

	"strings"
)

type IQueryUserPartnerCouponsResultV2 struct {

	// 优惠券ID。
	CouponId *string `json:"coupon_id,omitempty"`

	// 优惠券状态： 1：未激活2：可使用3：已使用4：已过期5：已回收
	Status *int32 `json:"status,omitempty"`

	// 客户账号ID。
	CustomerId *string `json:"customer_id,omitempty"`

	// 优惠券类别： 1：代金券4：现金券
	CouponType *int32 `json:"coupon_type,omitempty"`

	// 优惠券面额单位。 1：元。
	MeasureId *int32 `json:"measure_id,omitempty"`

	// 优惠券面值。
	FaceValue *decimal.Decimal `json:"face_value,omitempty"`

	// 生效时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。
	EffectiveTime *string `json:"effective_time,omitempty"`

	// 失效时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。
	ExpireTime *string `json:"expire_time,omitempty"`

	// 订单ID。
	OrderId *string `json:"order_id,omitempty"`

	// 促销计划ID。
	PromotionPlanId *string `json:"promotion_plan_id,omitempty"`

	// 促销计划名称。
	PromotionPlanName *string `json:"promotion_plan_name,omitempty"`

	// 促销计划描述。
	PromotionPlanDesc *string `json:"promotion_plan_desc,omitempty"`

	// 介质类型。 1：电子券2：纸质券
	MediaType *int32 `json:"media_type,omitempty"`

	// 获取方式。 1：线上领取2：线上兑换3：线上发放4：线下获取5：事件赠送
	FetchMethod *int32 `json:"fetch_method,omitempty"`

	// 优惠券限制。 具体请参见表3。
	UseLimits *[]ICouponUseLimitInfoV2 `json:"use_limits,omitempty"`

	// 优惠券的激活时间。
	ActiveTime *string `json:"active_time,omitempty"`

	// 优惠券的使用时间。
	LastUsedTime *string `json:"last_used_time,omitempty"`

	// 促销活动ID。
	PromotionId *string `json:"promotion_id,omitempty"`

	// 优惠券的创建时间。
	CreateTime *string `json:"create_time,omitempty"`

	// 优惠券余额。
	Balance *decimal.Decimal `json:"balance,omitempty"`

	// 锁定优惠券的订单ID。 如果为老版本优惠券，该值为空。
	LockOrderId *string `json:"lock_order_id,omitempty"`

	// 优惠券是否冻结。 0：否1：是 可用优惠券接口返回时不包括冻结状态的优惠券。
	IsFrozen *string `json:"is_frozen,omitempty"`
}

func (o IQueryUserPartnerCouponsResultV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IQueryUserPartnerCouponsResultV2 struct{}"
	}

	return strings.Join([]string{"IQueryUserPartnerCouponsResultV2", string(data)}, " ")
}
