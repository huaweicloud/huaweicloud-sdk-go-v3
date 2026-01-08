package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResourcePackagesOrderResponse Response Object
type CreateResourcePackagesOrderResponse struct {

	// 订单号，下单成功时返回订单ID。
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateResourcePackagesOrderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourcePackagesOrderResponse struct{}"
	}

	return strings.Join([]string{"CreateResourcePackagesOrderResponse", string(data)}, " ")
}
