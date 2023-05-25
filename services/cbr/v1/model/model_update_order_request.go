package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateOrderRequest struct {

	// 订单ID
	OrderId string `json:"order_id"`

	Body *CbcUpdate `json:"body,omitempty"`
}

func (o UpdateOrderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOrderRequest struct{}"
	}

	return strings.Join([]string{"UpdateOrderRequest", string(data)}, " ")
}
