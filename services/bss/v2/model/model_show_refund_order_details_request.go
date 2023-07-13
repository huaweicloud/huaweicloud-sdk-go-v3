package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRefundOrderDetailsRequest Request Object
type ShowRefundOrderDetailsRequest struct {

	// 退订订单或者降配订单的ID。
	OrderId string `json:"order_id"`
}

func (o ShowRefundOrderDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRefundOrderDetailsRequest struct{}"
	}

	return strings.Join([]string{"ShowRefundOrderDetailsRequest", string(data)}, " ")
}
