package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOrderCouponsByOrderIdRequest Request Object
type ListOrderCouponsByOrderIdRequest struct {

	// 订单ID。
	OrderId string `json:"order_id"`
}

func (o ListOrderCouponsByOrderIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOrderCouponsByOrderIdRequest struct{}"
	}

	return strings.Join([]string{"ListOrderCouponsByOrderIdRequest", string(data)}, " ")
}
