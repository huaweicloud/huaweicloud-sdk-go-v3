package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateChangeOrderResponse Response Object
type CreateChangeOrderResponse struct {

	// 订单号，下单成功时返回订单ID。
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateChangeOrderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateChangeOrderResponse struct{}"
	}

	return strings.Join([]string{"CreateChangeOrderResponse", string(data)}, " ")
}
