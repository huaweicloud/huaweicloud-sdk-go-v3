package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubscribePackageProductsResponse Response Object
type SubscribePackageProductsResponse struct {

	// 订购套餐包生成的订单号。
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SubscribePackageProductsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscribePackageProductsResponse struct{}"
	}

	return strings.Join([]string{"SubscribePackageProductsResponse", string(data)}, " ")
}
