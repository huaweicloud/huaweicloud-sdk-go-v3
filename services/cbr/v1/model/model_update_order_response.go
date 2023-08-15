package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateOrderResponse Response Object
type UpdateOrderResponse struct {

	// 订单ID
	OrderId *string `json:"orderId,omitempty"`

	// 变更状态码
	RetCode *string `json:"retCode,omitempty"`

	// 变更信息
	RetMsg         *string `json:"retMsg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateOrderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOrderResponse struct{}"
	}

	return strings.Join([]string{"UpdateOrderResponse", string(data)}, " ")
}
