package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IQueryUserCouponsResultV2 struct {

	// 优惠券实例ID。
	CouponId *string `json:"coupon_id,omitempty"`

	// 优惠券编码。
	CouponCode *string `json:"coupon_code,omitempty"`

	// 优惠券状态： 1：未激活2：待使用3：已使用4：已过期5：已回收
	Status *int32 `json:"status,omitempty"`

	// 客户账号ID。
	CustomerId *string `json:"customer_id,omitempty"`

	// 优惠券类型： 1：代金券2：折扣券3：产品券4：现金券
	CouponType *int32 `json:"coupon_type,omitempty"`

	// 度量单位。 1：元
	MeasureId *int32 `json:"measure_id,omitempty"`

	// 优惠券金额。
	FaceValue *float64 `json:"face_value,omitempty"`

	// 生效时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。
	ValidTime *string `json:"valid_time,omitempty"`

	// 失效时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。
	ExpireTime *string `json:"expire_time,omitempty"`

	// 订单ID。
	OrderId *string `json:"order_id,omitempty"`

	// 促销计划ID。
	PromotionPlanId *string `json:"promotion_plan_id,omitempty"`

	// 促销计划名称。
	PlanName *string `json:"plan_name,omitempty"`

	// 促销计划描述。
	PlanDesc *string `json:"plan_desc,omitempty"`

	// 介质类型。 1：电子券2：纸质券
	MediaType *int32 `json:"media_type,omitempty"`

	// 获取方式： 1：线上领取2：线上兑换3：线上发放4：线下获取5：事件赠送
	FetchMethod *int32 `json:"fetch_method,omitempty"`

	// 优惠券使用限制。 具体请参见表3。
	UseLimits *[]ICouponUseLimitInfoV2 `json:"use_limits,omitempty"`

	// 激活时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。
	ActiveTime *string `json:"active_time,omitempty"`

	// 使用时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。
	ReserveTime *string `json:"reserve_time,omitempty"`

	// 促销ID。
	PromotionId *string `json:"promotion_id,omitempty"`

	// 创建时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。
	CreateTime *string `json:"create_time,omitempty"`

	// 优惠券版本： 1：老版本，老版本优惠券只能使用一次2：新版本，新版本优惠券可以反复使用
	CouponVersion *int32 `json:"coupon_version,omitempty"`

	// 优惠券余额。单位：元。 如果为老版本优惠券，该值为空。
	Balance *float64 `json:"balance,omitempty"`

	// 锁定优惠券的订单ID。 如果为老版本优惠券，该值为空。
	LockOrderId *string `json:"lock_order_id,omitempty"`

	// 优惠券用途。
	CouponUsage *string `json:"coupon_usage,omitempty"`

	// 优惠券是否冻结： 0：否1：是
	IsFrozen *string `json:"is_frozen,omitempty"`

	// 币种。 CNY：人民币
	Currency *string `json:"currency,omitempty"`

	// 扩展字段。
	ExtendParam1 *string `json:"extend_param1,omitempty"`

	// 发券来源。 如果是合作伙伴发送的券，此处为伙伴ID。如果是活动发券，此处为活动ID：云豆兑换优惠券：云豆计划ID累计送优惠券：累计送计划ID抽奖送优惠券：抽奖计划ID事件送优惠券：事件计划ID定制优惠券：创建人ID
	SourceId *string `json:"source_id,omitempty"`
}

func (o IQueryUserCouponsResultV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IQueryUserCouponsResultV2 struct{}"
	}

	return strings.Join([]string{"IQueryUserCouponsResultV2", string(data)}, " ")
}
