package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CouponInfoV2 struct {

	// 优惠券实例ID。
	CouponId *string `json:"coupon_id,omitempty"`

	// 优惠券编码。
	CouponCode *string `json:"coupon_code,omitempty"`

	// 优惠券状态： 1：未激活2：待使用
	Status *int32 `json:"status,omitempty"`

	// 优惠券类型。 1：代金券2：折扣券3：产品券4：现金券
	CouponType *int32 `json:"coupon_type,omitempty"`

	// 面额单位： 1：元。
	MeasureId *int32 `json:"measure_id,omitempty"`

	// 面值。
	FaceValue *float64 `json:"face_value,omitempty"`

	// 生效时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。
	EffectiveTime *string `json:"effective_time,omitempty"`

	// 失效时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。
	ExpireTime *string `json:"expire_time,omitempty"`

	// 促销计划名称。
	PlanName *string `json:"plan_name,omitempty"`

	// 促销计划描述。
	PlanDesc *string `json:"plan_desc,omitempty"`

	// 优惠券限制。 具体请参见表3。
	UseLimits *[]LimitInfoV2 `json:"use_limits,omitempty"`

	// 激活时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。
	ActiveTime *string `json:"active_time,omitempty"`

	// 上一次使用时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。
	LastUsedTime *string `json:"last_used_time,omitempty"`

	// 创建时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。
	CreateTime *string `json:"create_time,omitempty"`

	// 优惠券版本。 1：老版本（包含三种：代金券、折扣券和奖金券）2：新版本（只有代金券）
	CouponVersion *int32 `json:"coupon_version,omitempty"`

	// 余额。
	Balance *float64 `json:"balance,omitempty"`

	// 使用优惠券的订单ID。表示有订单正在使用这个优惠券，优惠券已被锁定，只有锁定优惠券的订单才能使用这个优惠券，其他订单不能使用该优惠券。
	UsedByOrderId *string `json:"used_by_order_id,omitempty"`

	// 优惠券用途。
	CouponUsage *string `json:"coupon_usage,omitempty"`

	// 优惠券分组。 1：云商店发放的券2：华为云券-1024-专用代金券3：华为云券-使用限制-抵扣硬件的券0：华为云服务券（排除上述取值之外的券）
	CouponGroup *int32 `json:"coupon_group,omitempty"`
}

func (o CouponInfoV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CouponInfoV2 struct{}"
	}

	return strings.Join([]string{"CouponInfoV2", string(data)}, " ")
}
