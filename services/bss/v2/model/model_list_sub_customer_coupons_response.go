package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubCustomerCouponsResponse Response Object
type ListSubCustomerCouponsResponse struct {

	// 符合查询条件的总条数。
	Count *int32 `json:"count,omitempty"`

	// 优惠券记录。 具体请参见表2。
	UserCoupons    *[]IQueryUserCouponsResultV2 `json:"user_coupons,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListSubCustomerCouponsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubCustomerCouponsResponse struct{}"
	}

	return strings.Join([]string{"ListSubCustomerCouponsResponse", string(data)}, " ")
}
