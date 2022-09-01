package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 订单信息, 有订单表示包周期
type BillingInfoDict struct {

	// 订单信息
	OrderId *string `json:"order_id,omitempty" xml:"order_id"`

	// 产品id
	ProductId *string `json:"product_id,omitempty" xml:"product_id"`
}

func (o BillingInfoDict) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BillingInfoDict struct{}"
	}

	return strings.Join([]string{"BillingInfoDict", string(data)}, " ")
}
