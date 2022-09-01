package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CouponInfoV2 struct {

	// 优惠券实例ID。
	CouponId *string `json:"coupon_id,omitempty" xml:"coupon_id"`

	// 优惠券编码。
	CouponCode *string `json:"coupon_code,omitempty" xml:"coupon_code"`

	// 优惠券状态： 1：未激活2：待使用
	Status *int32 `json:"status,omitempty" xml:"status"`

	// 优惠券类型。 1：代金券2：折扣券3：产品券4：现金券
	CouponType *int32 `json:"coupon_type,omitempty" xml:"coupon_type"`

	// 面额单位： 1：元。
	MeasureId *int32 `json:"measure_id,omitempty" xml:"measure_id"`

	// 面值。
	FaceValue *float64 `json:"face_value,omitempty" xml:"face_value"`

	// 生效时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。
	EffectiveTime *string `json:"effective_time,omitempty" xml:"effective_time"`

	// 失效时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。
	ExpireTime *string `json:"expire_time,omitempty" xml:"expire_time"`

	// 促销计划名称。
	PlanName *string `json:"plan_name,omitempty" xml:"plan_name"`

	// 促销计划描述。
	PlanDesc *string `json:"plan_desc,omitempty" xml:"plan_desc"`

	// 优惠券限制。 具体请参见表3。
	UseLimits *[]LimitInfoV2 `json:"use_limits,omitempty" xml:"use_limits"`

	// 激活时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。
	ActiveTime *string `json:"active_time,omitempty" xml:"active_time"`

	// 上一次使用时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。
	LastUsedTime *string `json:"last_used_time,omitempty" xml:"last_used_time"`

	// 创建时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。
	CreateTime *string `json:"create_time,omitempty" xml:"create_time"`

	// 优惠券版本。 1：老版本（包含三种：代金券、折扣券和奖金券）2：新版本（只有代金券）
	CouponVersion *int32 `json:"coupon_version,omitempty" xml:"coupon_version"`

	// 余额。
	Balance *float64 `json:"balance,omitempty" xml:"balance"`

	// 使用优惠券的订单ID。表示有订单正在使用这个优惠券，优惠券已被锁定，只有锁定优惠券的订单才能使用这个优惠券，其他订单不能使用该优惠券。
	UsedByOrderId *string `json:"used_by_order_id,omitempty" xml:"used_by_order_id"`

	// 优惠券用途。
	CouponUsage *string `json:"coupon_usage,omitempty" xml:"coupon_usage"`

	// 优惠券分组。 1：云商店发放的券2：华为云券-1024-专用代金券3：华为云券-使用限制-抵扣硬件的券0：华为云服务券（排除上述取值之外的券）
	CouponGroup *int32 `json:"coupon_group,omitempty" xml:"coupon_group"`
}

func (o CouponInfoV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CouponInfoV2 struct{}"
	}

	return strings.Join([]string{"CouponInfoV2", string(data)}, " ")
}
