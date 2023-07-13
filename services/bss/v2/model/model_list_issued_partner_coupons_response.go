package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIssuedPartnerCouponsResponse Response Object
type ListIssuedPartnerCouponsResponse struct {

	// 总数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 发放的优惠券记录。 具体请参见表2。
	UserCoupons    *[]IQueryUserPartnerCouponsResultV2 `json:"user_coupons,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o ListIssuedPartnerCouponsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIssuedPartnerCouponsResponse struct{}"
	}

	return strings.Join([]string{"ListIssuedPartnerCouponsResponse", string(data)}, " ")
}
