package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OrderExtendParam 订单的扩展参数。
type OrderExtendParam struct {

	// 是否自动付款。true:自动支付; false：不自动支付;默认false。
	IsAutoPay *string `json:"is_auto_pay,omitempty"`
}

func (o OrderExtendParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrderExtendParam struct{}"
	}

	return strings.Join([]string{"OrderExtendParam", string(data)}, " ")
}
