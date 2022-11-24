package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CouponMaxUseQuantity struct {

	// 优惠券类型。 1：代金券2：折扣券3：产品券4：现金券
	CouponType *int32 `json:"coupon_type,omitempty"`

	// 优惠券分组。 1：云商店发放的券2：华为云券-1024-专用代金券3：华为云券-使用限制-抵扣硬件的券0：华为云服务券（排除上述取值之外的券）
	CouponGroup *int32 `json:"coupon_group,omitempty"`

	// 优惠券使用数量。
	UseQuantityValue *int32 `json:"use_quantity_value,omitempty"`
}

func (o CouponMaxUseQuantity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CouponMaxUseQuantity struct{}"
	}

	return strings.Join([]string{"CouponMaxUseQuantity", string(data)}, " ")
}
