package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListOrderCouponsByOrderIdResponse struct {
	// 查询总数。

	Count *int32 `json:"count,omitempty"`
	// 可用的优惠券列表。 具体请参见表2。

	UserCoupons    *[]CouponInfoV2 `json:"user_coupons,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListOrderCouponsByOrderIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOrderCouponsByOrderIdResponse struct{}"
	}

	return strings.Join([]string{"ListOrderCouponsByOrderIdResponse", string(data)}, " ")
}
