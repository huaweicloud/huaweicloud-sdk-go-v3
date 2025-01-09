package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDesktopOrderResponse Response Object
type CreateDesktopOrderResponse struct {

	// 订单号，下单成功时返回订单ID。
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDesktopOrderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDesktopOrderResponse struct{}"
	}

	return strings.Join([]string{"CreateDesktopOrderResponse", string(data)}, " ")
}
