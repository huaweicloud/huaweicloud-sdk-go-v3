package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResizeOrderResponse Response Object
type CreateResizeOrderResponse struct {

	// 订单ID，仅在创建包周期实例时返回。
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateResizeOrderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResizeOrderResponse struct{}"
	}

	return strings.Join([]string{"CreateResizeOrderResponse", string(data)}, " ")
}
