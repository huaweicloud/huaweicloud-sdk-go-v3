package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceOrderResponse Response Object
type CreateInstanceOrderResponse struct {

	// 订单ID。
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateInstanceOrderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceOrderResponse struct{}"
	}

	return strings.Join([]string{"CreateInstanceOrderResponse", string(data)}, " ")
}
