package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateOrderFormResponse struct {

	// 订单ID
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateOrderFormResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrderFormResponse struct{}"
	}

	return strings.Join([]string{"CreateOrderFormResponse", string(data)}, " ")
}
