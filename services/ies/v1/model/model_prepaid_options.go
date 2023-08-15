package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PrepaidOptions 包年/包月计费信息
type PrepaidOptions struct {

	// 订单ID
	OrderId *string `json:"order_id,omitempty"`

	// 产品ID
	ProductId *string `json:"product_id,omitempty"`

	PayMode *PayMode `json:"pay_mode,omitempty"`
}

func (o PrepaidOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrepaidOptions struct{}"
	}

	return strings.Join([]string{"PrepaidOptions", string(data)}, " ")
}
