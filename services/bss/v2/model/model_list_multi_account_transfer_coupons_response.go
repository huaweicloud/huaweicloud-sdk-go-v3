package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMultiAccountTransferCouponsResponse Response Object
type ListMultiAccountTransferCouponsResponse struct {

	// 记录条数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 可拨款优惠券记录。 具体请参见表2。
	AvailTransferCoupons *[]AvailTransferCoupon `json:"avail_transfer_coupons,omitempty"`
	HttpStatusCode       int                    `json:"-"`
}

func (o ListMultiAccountTransferCouponsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMultiAccountTransferCouponsResponse struct{}"
	}

	return strings.Join([]string{"ListMultiAccountTransferCouponsResponse", string(data)}, " ")
}
