package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AvailRetrieveCoupons struct {

	// 主优惠券ID。
	CouponId *string `json:"coupon_id,omitempty"`

	// 促销计划名称。
	PlanName *string `json:"plan_name,omitempty"`

	// 子优惠券ID。主优惠券拨款后生成的子优惠券ID。
	SubCouponId *string `json:"sub_coupon_id,omitempty"`

	// 优惠券余额。单位为元
	Balance *string `json:"balance,omitempty"`

	// 生效时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。
	EffectiveTime *string `json:"effective_time,omitempty"`

	// 失效时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。
	ExpireTime *string `json:"expire_time,omitempty"`

	// 优惠券限制。 具体请参见表3。
	UseLimits *[]ICouponUseLimitInfoV2 `json:"use_limits,omitempty"`
}

func (o AvailRetrieveCoupons) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AvailRetrieveCoupons struct{}"
	}

	return strings.Join([]string{"AvailRetrieveCoupons", string(data)}, " ")
}
