package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateProductOrderResponse Response Object
type CreateProductOrderResponse struct {

	// 订单ID
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateProductOrderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProductOrderResponse struct{}"
	}

	return strings.Join([]string{"CreateProductOrderResponse", string(data)}, " ")
}
